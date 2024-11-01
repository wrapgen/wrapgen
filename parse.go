// Copyright 2024 Wrapgen authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/fs"
	"log/slog"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/wrapgen/wrapgen/internal/argparse"
	"github.com/wrapgen/wrapgen/internal/cache"
	"github.com/wrapgen/wrapgen/internal/modinfo"
)

type packageSpec struct {
	modulePath, packagePath, srcDir string
}

type parseContext struct {
	directories    *cache.Cache[packageSpec, *directoryParser]
	modinfoLoader  map[string]*modinfo.Loader // [module-path>]:*loader
	dirPackageSpec *cache.Cache[string, packageSpec]
}

func newParseContext() *parseContext {
	return &parseContext{
		directories:    cache.New[packageSpec, *directoryParser](),
		modinfoLoader:  map[string]*modinfo.Loader{},
		dirPackageSpec: cache.New[string, packageSpec](),
	}
}

type directoryParser struct {
	packagePath   string
	fileSet       *token.FileSet
	pkgs          []*packageParser
	srcDir        string
	modinfoLoader *modinfo.Loader
	modulePath    string
}

type packageParser struct {
	directoryParser *directoryParser
	pkg             *ast.Package
	parsed          bool
	relevant        bool
	files           []*fileParser
}

type fileParser struct {
	packageParser       *packageParser
	file                *ast.File
	Imports             map[string]string // <package-path> - <alias, is often empty>
	InterfaceGenerators []*interfaceGenerator
	Interface           []*Interface
	packageMap          map[string]string
}

type interfaceGenerator struct {
	interfaceParser *Interface
	PackageName     string
	Template        string
	Destination     string
	Vars            map[string]any
	Name            string
	Position        string
}

// ResolveImportSelectorToPackagePath takes package name as argument like "redis",
// then checks if any import of this file has an alias named "redis", if so returns that packagePath.
// Otherwise, it checks the global package map and returns package Path from there.
// If not found it returns "".
func (fp *fileParser) ResolveImportSelectorToPackagePath(name string) string {
	for packagePath, importAlias := range fp.Imports {
		if importAlias == name {
			return packagePath
		}
	}

	for packagePath, packageName := range fp.packageMap {
		if packageName == name {
			return packagePath
		}
	}

	return ``
}

func (p *parseContext) readPackage(ps packageSpec) (*directoryParser, error) {
	slog.Info("readPackageInDirectory", "packagePath", ps.packagePath)

	dirP := &directoryParser{
		packagePath:   ps.packagePath,
		fileSet:       token.NewFileSet(),
		pkgs:          make([]*packageParser, 0),
		srcDir:        ps.srcDir,
		modinfoLoader: modinfo.NewLoader(ps.srcDir),
		modulePath:    ps.modulePath,
	}

	srcDirAbs, err := filepath.Abs(ps.srcDir)
	if err != nil {
		return nil, err
	}
	b := build.Default
	b.Dir = srcDirAbs
	var pkgs map[string]*ast.Package
	if imp, err := b.Import(ps.packagePath, srcDirAbs, 0); err != nil {
		return nil, err
	} else if pkgs, err = parser.ParseDir(dirP.fileSet, imp.Dir, func(info fs.FileInfo) bool {
		res := slices.Contains(imp.GoFiles, info.Name())
		// for the local module, also parse the _test.go files.
		if strings.HasPrefix(srcDirAbs, imp.Root) {
			res = res || slices.Contains(imp.TestGoFiles, info.Name())
		}
		return res
	}, parser.ParseComments); err != nil {
		return nil, err
	}

	for _, p := range pkgs {
		dirP.pkgs = append(dirP.pkgs, &packageParser{
			directoryParser: dirP,
			pkg:             p,
		})
	}

	// collect import paths to reduce calls to 'go list'.
	var allImportPaths []string
	for _, pp := range dirP.pkgs {
		for _, file := range pp.pkg.Files {
			fp := &fileParser{
				packageParser: pp,
				file:          file,
				Imports:       make(map[string]string),
			}
			pp.files = append(pp.files, fp)

			for _, is := range file.Imports {
				importPath := is.Path.Value[1 : len(is.Path.Value)-1] // remove quotes
				if is.Name != nil {
					fp.Imports[importPath] = is.Name.Name
				} else {
					fp.Imports[importPath] = ""
				}
			}
			for packagePath := range fp.Imports {
				allImportPaths = append(allImportPaths, packagePath)
			}
		}
	}

	allImportPaths = append(allImportPaths, dirP.packagePath)
	_, err = dirP.modinfoLoader.PackageMap(allImportPaths)
	if err != nil {
		return nil, err
	}

	for _, pp := range dirP.pkgs {
		for _, fp := range pp.files {
			// resolve package-import-paths to package names.
			// For example packagePath "bla.com/client/v1" is usually named "client".
			packageMap, err := dirP.modinfoLoader.PackageMap(append(keys(fp.Imports), fp.packageParser.directoryParser.packagePath))
			if err != nil {
				return nil, err
			}
			fp.packageMap = packageMap
		}

		err := pp.readInterfaces(p)
		if err != nil {
			return nil, err
		}
	}

	return dirP, nil
}

func (p *parseContext) GetPackageSpec(inputFilePath string) (packageSpec, error) {
	ps, err := p.dirPackageSpec.GetOrAdd(filepath.Dir(inputFilePath), func(dir string) (packageSpec, error) {
		pp, mp, err := modinfo.ImportPath(dir)
		if err != nil {
			return packageSpec{}, err
		}
		return packageSpec{
			modulePath:  mp,
			packagePath: pp,
			srcDir:      dir,
		}, nil
	})
	if err != nil {
		return packageSpec{}, err
	}
	return ps, nil
}

func (p *parseContext) ReadPackageByContainedFile(inputFilePath string) (*directoryParser, error) {
	ps, err := p.GetPackageSpec(inputFilePath)
	if err != nil {
		return nil, err
	}

	return p.ReadPackage(ps)
}

func (p *parseContext) ReadPackage(ps packageSpec) (*directoryParser, error) {
	return p.directories.GetOrAdd(ps, p.readPackage)
}

func (pp *packageParser) position(pos token.Pos) string {
	ps := pp.directoryParser.fileSet.Position(pos)
	args := []any{ps.Filename, ps.Line, ps.Column}
	return fmt.Sprintf("%s:%d:%d: ", args...)
}

func (pp *packageParser) errorf(pos token.Pos, format string, args ...any) error {
	return fmt.Errorf("%s "+format, append([]any{pp.position(pos)}, args...)...)
}

func (pp *packageParser) readInterfaces(p *parseContext) error {
	if pp.parsed == true {
		return nil
	}
	pp.parsed = true

	// first run to find all interfaces in all files.
	for _, fp := range pp.files {
		for _, decl := range fp.file.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				it, ok := ts.Type.(*ast.InterfaceType)
				if !ok {
					continue
				}

				ip := &Interface{
					fileParser: fp,
					Interface:  it,
					typeSpec:   ts,
					tps:        make(map[string]Type),
					Name:       ts.Name.Name,
					Package:    fp.packageParser.directoryParser.packagePath,
				}

				var comments []*ast.Comment
				if gd.Doc != nil {
					comments = gd.Doc.List
				}

				if ts, ok := spec.(*ast.TypeSpec); ok {
					if ts.Doc != nil {
						comments = append(comments, ts.Doc.List...)
					}
				}

				for _, comment := range comments {
					if commandLine, found := strings.CutPrefix(comment.Text, wrapgenGenerateKeyword); found {
						cmd := &interfaceGenerator{
							interfaceParser: ip,
						}
						f := flag.NewFlagSet("wrapgen:generate", flag.PanicOnError)
						f.StringVar(&cmd.PackageName, "package", "", "Target package. Default: basename of destination.")
						f.StringVar(&cmd.Template, "template", "", "Builtin template or path to file if it contains /.")
						f.StringVar(&cmd.Destination, "destination", "", "Destination file.")
						f.StringVar(&cmd.Name, "name", ip.Name, "Interface name to assume.")
						vars := f.String("vars", "", "Template variable in var1=value1,var2=value2 style.")

						args, err := argparse.Split(commandLine)
						if err != nil {
							return fp.packageParser.errorf(comment.Pos(), "reading argument list: %w", err)
						}
						if f.NArg() > 0 {
							return fp.packageParser.errorf(comment.Pos(), "unexpected args: %v", f.Args())
						}

						err = f.Parse(args)
						if err != nil {
							return fp.packageParser.errorf(comment.Pos(), "parsing flags failed")
						}

						cmd.Vars = make(map[string]any)
						for _, pair := range strings.Split(*vars, ",") {
							if pair == "" {
								continue
							}
							kv := strings.Split(pair, "=")
							if len(kv) == 2 {
								cmd.Vars[kv[0]] = kv[1]
							}
						}

						if cmd.Destination == "" {
							return fp.packageParser.errorf(comment.Pos(), "-destination must not be empty")
						}

						// convert slashes to native host format.
						cmd.Destination = filepath.FromSlash(cmd.Destination)
						// convert to path relative to source directory of the comment.
						cmd.Destination = filepath.Join(fp.packageParser.directoryParser.srcDir, cmd.Destination)

						if cmd.PackageName == "" {
							// this is a simplification. Ideally we should check the packageMap.
							destAbs, err := filepath.Abs(cmd.Destination)
							if err != nil {
								return err
							}
							cmd.PackageName = filepath.Base(filepath.Dir(destAbs))
						}

						cmd.Position = pp.directoryParser.fileSet.Position(comment.Pos()).String()
						fp.InterfaceGenerators = append(fp.InterfaceGenerators, cmd)
					}
				}

				fp.Interface = append(fp.Interface, ip)
			}
		}
	}

	// second run when all interfaces have been created but not yet initialized.
	// This only descends into interfaces with a //wrapgen:generate annotation and
	// recursively parses all embedded interfaces.
	for _, fp := range pp.files {
		for _, ig := range fp.InterfaceGenerators {
			err := ig.interfaceParser.parse(p)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ip *Interface) parse(p *parseContext) error {
	ip.lock.Lock()
	defer ip.lock.Unlock()
	if ip.parsed {
		return nil
	}

	slog.Info("parse", "interface", ip.fileParser.packageParser.directoryParser.packagePath+"."+ip.Name)

	pp := ip.fileParser.packageParser

	if ip.typeSpec.TypeParams != nil {
		for _, tp := range ip.typeSpec.TypeParams.List {
			for _, tm := range tp.Names {
				switch v := tm.Obj.Decl.(type) {
				case *ast.Field:
					ip.tps[tm.Name] = nil

					t, err := p.parseType(ip.fileParser, v.Type, nil)
					if err != nil {
						return err
					}

					ip.TypeParams = append(ip.TypeParams, &Parameter{
						Name: tm.Name,
						Type: t,
					})
				default:
					return ip.fileParser.packageParser.errorf(tp.Pos(), "fail to parse generic type parameter")
				}
			}
		}
	}

	for _, field := range ip.Interface.Methods.List {
		switch v := field.Type.(type) {
		case *ast.FuncType: // Regular function of the interface.
			_ = v
			if nn := len(field.Names); nn != 1 {
				return pp.errorf(v.Pos(),
					"expected one name for interface %v, got %d", ip.typeSpec.Name.Name, nn)
			}
			m := &Method{
				Name:      field.Names[0].String(),
				Interface: ip,
			}
			var err error
			m.In, m.Variadic, m.Out, err = p.parseFunc(ip.fileParser, v, ip.tps)
			if err != nil {
				return err
			}
			ip.Methods = append(ip.Methods, m)
		case *ast.Ident: // Embedded interface in same package or stdlib interface.
			if v.Name == "error" {
				ip.Embedded = append(ip.Embedded, ErrorInterface)
			} else {
				var ipTargetInterface *Interface
			findIntf:
				for _, fp := range pp.files {
					for _, ip := range fp.Interface {
						if ip.Name == v.Name {
							ipTargetInterface = ip
							break findIntf
						}
					}
				}
				if ipTargetInterface == nil {
					slog.Warn("TODO, embed", "name", v.Name,
						"position", ip.fileParser.packageParser.position(v.Pos()))
				} else {
					err := ipTargetInterface.parse(p)
					if err != nil {
						return err
					}
					ip.Embedded = append(ip.Embedded, ipTargetInterface)
				}
			}
		case *ast.SelectorExpr: // Embedded interface in another package.
			importName := v.X.(*ast.Ident).Name
			interfaceName := v.Sel.Name

			packagePath := ip.fileParser.ResolveImportSelectorToPackagePath(importName)
			if packagePath == "" {
				return pp.errorf(v.Pos(), "could not resolve import path")
			}
			dp, err := p.ReadPackage(packageSpec{
				modulePath:  ip.fileParser.packageParser.directoryParser.modulePath,
				packagePath: packagePath,
				srcDir:      ip.fileParser.packageParser.directoryParser.srcDir,
			})
			if err != nil {
				return err
			}
			var ppTargetInterface *packageParser
			for _, pkg := range dp.pkgs {
				if pkg.pkg.Name == ip.fileParser.packageMap[packagePath] {
					ppTargetInterface = pkg
					break
				}
			}
			if ppTargetInterface == nil {
				return pp.errorf(v.Pos(), "could not resolve import package")
			}

			var ipTargetInterface *Interface
		findIP:
			for _, fp := range ppTargetInterface.files {
				for _, ip := range fp.Interface {
					if ip.Name == interfaceName {
						ipTargetInterface = ip
						break findIP
					}
				}
			}
			if ipTargetInterface == nil {
				return pp.errorf(v.Pos(),
					"could not resolve embedded interface %vin %v", interfaceName, packagePath)
			}
			err = ipTargetInterface.parse(p)
			if err != nil {
				return err
			}
			ip.Embedded = append(ip.Embedded, ipTargetInterface)
		default:
			slog.Warn("Unhandled interface field", "position", pp.position(v.Pos()))
		}
	}

	// Recursively collect all methods of embedded interfaces.
	ip.AllMethods = make([]*Method, 0, len(ip.Methods))
	var appendAllMethods func(i *Interface)
	appendAllMethods = func(i *Interface) {
		ip.AllMethods = append(ip.AllMethods, i.Methods...)
		for _, embedded := range i.Embedded {
			appendAllMethods(embedded)
		}
	}
	appendAllMethods(ip)
	sort.Slice(ip.AllMethods, func(i, j int) bool {
		return ip.AllMethods[i].Name < ip.AllMethods[j].Name
	})

	// Eliminate duplicate methods, fail if same name but different signature.
	for i := 0; i < len(ip.AllMethods)-1; i++ {
		current, next := ip.AllMethods[i], ip.AllMethods[i+1]
		if current.Name == next.Name {
			if !reflect.DeepEqual(current.In, next.In) ||
				!reflect.DeepEqual(current.Out, next.Out) ||
				!reflect.DeepEqual(current.Variadic, next.Variadic) {
				return pp.errorf(ip.Interface.Pos(),
					"conflicting methods named %v on interface", current.Name)
			}
			ip.AllMethods = append(ip.AllMethods[:i], ip.AllMethods[i+1:]...)
		}
	}

	ip.parsed = true
	return nil
}

func (p *parseContext) parseFunc(fp *fileParser, f *ast.FuncType, tps map[string]Type) (inParam []*Parameter, variadic *Parameter, outParam []*Parameter, err error) {
	if f.Params != nil {
		regParams := f.Params.List
		if isVariadic(f) {
			n := len(regParams)
			varParams := regParams[n-1:]
			regParams = regParams[:n-1]
			vp, err := p.parseFieldList(fp, varParams, tps)
			if err != nil {
				return nil, nil, nil, fp.packageParser.errorf(varParams[0].Pos(), "failed parsing variadic argument: %v", err)
			}
			variadic = vp[0]
		}
		inParam, err = p.parseFieldList(fp, regParams, tps)
		if err != nil {
			return nil, nil, nil, fp.packageParser.errorf(f.Pos(), "failed parsing arguments: %v", err)
		}
	}
	if f.Results != nil {
		outParam, err = p.parseFieldList(fp, f.Results.List, tps)
		if err != nil {
			return nil, nil, nil, fp.packageParser.errorf(f.Pos(), "failed parsing returns: %v", err)
		}
	}
	return inParam, variadic, outParam, err
}

func (p *parseContext) parseFieldList(fp *fileParser, fields []*ast.Field, tps map[string]Type) ([]*Parameter, error) {
	nf := 0
	for _, f := range fields {
		nn := len(f.Names)
		if nn == 0 {
			nn = 1 // anonymous parameter
		}
		nf += nn
	}
	if nf == 0 {
		return nil, nil
	}
	ps := make([]*Parameter, nf)
	i := 0 // destination index
	for _, f := range fields {
		t, err := p.parseType(fp, f.Type, tps)
		if err != nil {
			return nil, err
		}

		if len(f.Names) == 0 {
			// anonymous arg
			ps[i] = &Parameter{Type: t}
			i++
			continue
		}
		for _, name := range f.Names {
			ps[i] = &Parameter{Name: name.Name, Type: t}
			i++
		}
	}
	return ps, nil
}

func (p *parseContext) parseType(fp *fileParser, typ ast.Expr, tps map[string]Type) (Type, error) {
	switch v := typ.(type) {
	case *ast.ArrayType:
		ln := -1
		if v.Len != nil {
			value, err := parseArrayLength(fp.packageParser, v.Len)
			if err != nil {
				return nil, err
			}
			ln, err = strconv.Atoi(value)
			if err != nil {
				return nil, fp.packageParser.errorf(v.Len.Pos(), "bad array size: %v", err)
			}
		}
		t, err := p.parseType(fp, v.Elt, tps)
		if err != nil {
			return nil, err
		}
		return &ArrayType{Len: ln, Type: t}, nil
	case *ast.ChanType:
		t, err := p.parseType(fp, v.Value, tps)
		if err != nil {
			return nil, err
		}
		var dir ChanDir
		if v.Dir == ast.SEND {
			dir = SendDir
		}
		if v.Dir == ast.RECV {
			dir = RecvDir
		}
		return &ChanType{Dir: dir, Type: t}, nil
	case *ast.Ellipsis:
		// assume we're parsing a variadic argument
		return p.parseType(fp, v.Elt, tps)
	case *ast.FuncType:
		in, variadic, out, err := p.parseFunc(fp, v, tps)
		if err != nil {
			return nil, err
		}
		return &FuncType{In: in, Out: out, Variadic: variadic}, nil
	case *ast.Ident:
		it, ok := tps[v.Name]
		if v.IsExported() && !ok {
			return &NamedType{
				fileParser: fp,
				Package:    fp.packageParser.directoryParser.packagePath,
				Type:       v.Name,
			}, nil
		}
		if ok && it != nil {
			return it, nil
		}
		if v.Obj == nil {
			// assume predeclared type
			return PredeclaredType(v.Name), nil
		}
		var typeParams []Type
		typeSpec, ok := v.Obj.Decl.(*ast.TypeSpec)
		if typeSpec != nil && typeSpec.TypeParams != nil {
			var err error
			params, err := p.parseFieldList(fp, typeSpec.TypeParams.List, tps)
			if err != nil {
				return nil, err
			}
			for _, p := range params {
				typeParams = append(typeParams, p.Type)
			}
			// assume predeclared or local type
			return &NamedType{
				fileParser: fp,
				Package:    fp.packageParser.directoryParser.packagePath,
				Type:       v.Name,
				TypeParams: &TypeParametersType{TypeParameters: typeParams},
			}, nil
		}
		// assume predeclared or local type
		return &NamedType{
			fileParser: fp,
			Package:    fp.packageParser.directoryParser.packagePath,
			Type:       v.Name,
		}, nil
	case *ast.InterfaceType:
		if v.Methods != nil && len(v.Methods.List) > 0 {
			return nil, fp.packageParser.errorf(v.Pos(), "can't handle non-empty unnamed interface types")
		}
		return PredeclaredType("any"), nil
	case *ast.MapType:
		key, err := p.parseType(fp, v.Key, tps)
		if err != nil {
			return nil, err
		}
		value, err := p.parseType(fp, v.Value, tps)
		if err != nil {
			return nil, err
		}
		return &MapType{Key: key, Value: value}, nil
	case *ast.SelectorExpr:
		pkgName := v.X.(*ast.Ident).String()

		for packagePath, importAlias := range fp.Imports {
			if importAlias == pkgName {
				return &NamedType{Package: packagePath, Type: v.Sel.String()}, nil
			}
			if importAlias == "" {
				if packageName, found := fp.packageMap[packagePath]; found {
					if packageName == pkgName {
						return &NamedType{Package: packagePath, Type: v.Sel.String()}, nil
					}
				}
			}
		}
		return nil, fp.packageParser.errorf(v.Pos(), "unknown package %q", pkgName)
	case *ast.StarExpr:
		t, err := p.parseType(fp, v.X, tps)
		if err != nil {
			return nil, err
		}
		return &PointerType{Type: t}, nil
	case *ast.StructType:
		if v.Fields != nil && len(v.Fields.List) > 0 {
			return nil, fp.packageParser.errorf(v.Pos(), "can't handle non-empty unnamed struct types")
		}
		return PredeclaredType("struct{}"), nil
	case *ast.ParenExpr:
		return p.parseType(fp, v.X, tps)
	case *ast.IndexExpr:
		m, err := p.parseType(fp, v.X, tps)
		if err != nil {
			return nil, err
		}
		nm, ok := m.(*NamedType)
		if !ok {
			return m, nil
		}
		t, err := p.parseType(fp, v.Index, tps)
		if err != nil {
			return nil, err
		}
		nm.TypeParams = &TypeParametersType{TypeParameters: []Type{t}}
		return m, nil
	case *ast.IndexListExpr:
		m, err := p.parseType(fp, v.X, tps)
		if err != nil {
			return nil, err
		}
		nm, ok := m.(*NamedType)
		if !ok {
			return m, nil
		}
		var ts []Type
		for _, expr := range v.Indices {
			t, err := p.parseType(fp, expr, tps)
			if err != nil {
				return nil, err
			}
			ts = append(ts, t)
		}
		nm.TypeParams = &TypeParametersType{TypeParameters: ts}
		return m, nil
	}

	return nil, fp.packageParser.errorf(typ.Pos(), "don't know how to parse type %T", typ)
}

func parseArrayLength(pp *packageParser, expr ast.Expr) (string, error) {
	switch val := expr.(type) {
	case (*ast.BasicLit):
		return val.Value, nil
	case (*ast.Ident):
		// when the length is a const defined locally
		return val.Obj.Decl.(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value, nil
	case (*ast.SelectorExpr):
		// when the length is a const defined in an external package
		usedPkg, err := importer.Default().Import(fmt.Sprintf("%s", val.X))
		if err != nil {
			return "", pp.errorf(expr.Pos(), "unknown package in array length: %v", err)
		}
		ev, err := types.Eval(token.NewFileSet(), usedPkg, token.NoPos, val.Sel.Name)
		if err != nil {
			return "", pp.errorf(expr.Pos(), "unknown constant in array length: %v", err)
		}
		return ev.Value.String(), nil
	case (*ast.ParenExpr):
		return parseArrayLength(pp, val.X)
	case (*ast.BinaryExpr):
		x, err := parseArrayLength(pp, val.X)
		if err != nil {
			return "", err
		}
		y, err := parseArrayLength(pp, val.Y)
		if err != nil {
			return "", err
		}
		biExpr := fmt.Sprintf("%s%v%s", x, val.Op, y)
		tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, biExpr)
		if err != nil {
			return "", pp.errorf(expr.Pos(), "invalid expression in array length: %v", err)
		}
		return tv.Value.String(), nil
	default:
		return "", pp.errorf(expr.Pos(), "invalid expression in array length: %v", val)
	}
}

// isVariadic returns whether the function is variadic.
func isVariadic(f *ast.FuncType) bool {
	nargs := len(f.Params.List)
	if nargs == 0 {
		return false
	}
	_, ok := f.Params.List[nargs-1].Type.(*ast.Ellipsis)
	return ok
}

// keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
