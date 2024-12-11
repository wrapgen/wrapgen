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

// WrapGen generates code based on Go interfaces using text/template.
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"go/build"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"sync"

	"github.com/wrapgen/wrapgen/internal/modinfo"
)

var (
	flagShowVersion = flag.Bool("version", false, "Print version.")
	flagVerbose     = flag.Bool("verbose", false, "Verbose logging.")
)

// tagsFlag is the implementation of the -tags flag.
// copied from go/internal/work/build.go
type tagsFlag []string

func (v *tagsFlag) Set(s string) error {
	// Split on commas, ignore empty strings.
	*v = []string{}
	for _, s := range strings.Split(s, ",") {
		if s != "" {
			*v = append(*v, s)
		}
	}
	return nil
}

func (v *tagsFlag) String() string {
	return "<TagsFlag>"
}

func init() {
	flag.Var((*tagsFlag)(&build.Default.BuildTags), "tags",
		"build tags like for 'go build' (sets go/build.Default.BuildTags)")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelWarn,
	}
	if *flagVerbose {
		opts.Level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &opts)))

	debuginit()
	defer debugfinish()

	if *flagShowVersion {
		printVersion()
		return
	}

	var args = flag.Args()
	if len(args) == 0 {
		args = append(args, ".")
	}

	processDirectories(func(path string) (stringWriteCloser, error) {
		return fsWriter(path)
	}, args...)
}

type stringWriteCloser interface {
	io.WriteCloser
	io.StringWriter
}

type fileWriter func(path string) (stringWriteCloser, error)

func processDirectories(w fileWriter, basePaths ...string) {
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, runtime.GOMAXPROCS(0))
	ctx, cancel := context.WithCancelCause(context.Background())

	pc, err := newParseContext(basePaths)
	if err != nil {
		slog.Error("initializing parseContext", "error", err)
		os.Exit(1)
	}

	err = pc.parsePaths()
	if err != nil {
		slog.Error("parse sources", "error", err)
		os.Exit(1)
	}

	// group generate commands per output file
	type fileFileGenerator struct {
		packageName string
		packagePath string
		interfaces  []*interfaceGenerator
		fileParser  *fileParser
	}
	commandsPerTargetFile := make(map[string]*fileFileGenerator)
	for _, dp := range pc.directories.Values() {
		for _, pp := range dp.pkgs {
			for _, fp := range pp.files {
				for _, generateCommand := range fp.InterfaceGenerators {
					targetFilePath := generateCommand.Destination
					if _, found := commandsPerTargetFile[targetFilePath]; !found {
						commandsPerTargetFile[targetFilePath] = new(fileFileGenerator)
						commandsPerTargetFile[targetFilePath].fileParser = generateCommand.interfaceParser.fileParser
					}
					commandsPerTargetFile[targetFilePath].interfaces = append(commandsPerTargetFile[targetFilePath].interfaces, generateCommand)
				}
			}
		}
	}

	// verify there are no contradicting targets in the output.
	for targetFilePath, fg := range commandsPerTargetFile {
		for _, cmd := range fg.interfaces {
			if fg.packageName == "" {
				fg.packageName = cmd.PackageName
				ps, err := pc.GetPackageSpec(targetFilePath)
				if err != nil {
					slog.Error("failed to process a file", "position", cmd.Position, "error", err)
					os.Exit(1)
				}
				fg.packagePath = ps.packagePath
			} else if fg.packageName != cmd.PackageName {
				slog.Error("contradicting target packages", "position", cmd.Position,
					"packageA", fg.packageName, "packageB", cmd.PackageName)
				os.Exit(1)
			}
		}
	}

	// process all target files in parallel.
	for targetFilePath, fg := range commandsPerTargetFile {
		select {
		case sem <- struct{}{}:
		case <-ctx.Done():
			break // a go-routine failed
		}
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-sem
			}()

			output := &outputBuffer{
				buf:                &bytes.Buffer{},
				packageMap:         make(map[string]string),
				packageName:        fg.packageName,
				packagePath:        fg.packagePath,
				importedPackages:   make(map[string]string),
				importPackageNames: make(map[string]string),
				modinfoLoader:      fg.fileParser.packageParser.directoryParser.modinfoLoader,
			}

			/// recursively merge the package map of all involved interfaces, also the embedded ones.
			var intfs []*Interface
			for _, ig := range fg.interfaces {
				intfs = append(intfs, ig.interfaceParser)
			}
			for i := 0; i < len(intfs); i++ {
				intfs = append(intfs, intfs[i].Embedded...)
				for packagePath, packageName := range intfs[i].fileParser.packageMap {
					output.packageMap[packagePath] = packageName
				}
			}

			for _, cmd := range fg.interfaces {
				if output.buf.Len() > 0 {
					output.buf.WriteRune('\n')
				}

				templateBasedir := cmd.interfaceParser.fileParser.packageParser.directoryParser.srcDir
				t, err := getTemplate(templateBasedir, cmd.Template)
				if err != nil {
					cancel(fmt.Errorf("%s: %s", cmd.Position, err))
					return
				}

				t, err = t.Clone()
				if err != nil {
					cancel(fmt.Errorf("clone template %s: %w", cmd.Position, err))
					return
				}
				t = t.Funcs(templateFunctions(output))

				err = t.Execute(output.buf, struct {
					Interface *Interface
					Vars      map[string]any
					Name      string
				}{
					Interface: cmd.interfaceParser,
					Vars:      cmd.Vars,
					Name:      cmd.Name,
				})
				if err != nil {
					slog.Error("failed to execute template", "template", cmd.Template,
						"position", cmd.Position, "error", err)
					os.Exit(1)
				}
			}

			// Sort imports into 3 groups
			var importGroups [3][]string
			for k := range output.importedPackages {
				if !strings.Contains(k, ".") {
					importGroups[0] = append(importGroups[0], k)
				} else if strings.HasPrefix(k, fg.fileParser.packageParser.directoryParser.modulePath) {
					importGroups[1] = append(importGroups[1], k)
				} else {
					importGroups[2] = append(importGroups[2], k)
				}
			}
			sort.Strings(importGroups[0])
			sort.Strings(importGroups[1])
			sort.Strings(importGroups[2])

			outputFile, err := w(targetFilePath)
			if err != nil {
				slog.Error("opening output file failed", "target", targetFilePath, "error", err)
				os.Exit(1)
			}
			defer outputFile.Close()

			_, _ = outputFile.WriteString("// Code generated by WrapGen. DO NOT EDIT.\n")
			_, _ = outputFile.WriteString("package " + fg.packageName + "\n\n")
			_, _ = outputFile.WriteString("import (\n")

			for i, group := range importGroups {
				if len(group) > 0 {
					if i > 0 {
						_, _ = outputFile.WriteString("\n")
					}
					for _, packagePath := range group {
						alias := output.importedPackages[packagePath]
						if alias != output.packageMap[packagePath] {
							_, _ = outputFile.WriteString("\t" + alias + ` "` + packagePath + `"` + "\n")
						} else {
							_, _ = outputFile.WriteString("\t" + `"` + packagePath + `"` + "\n")
						}
					}
				}
			}
			_, _ = outputFile.WriteString(")\n\n")
			_, _ = output.buf.WriteTo(outputFile)
		}()
	}

	wg.Wait()
	if err := ctx.Err(); err != nil {
		slog.Error("writing an output file failed", "error", context.Cause(ctx))
		os.Exit(1)
	}
}

func fsWriter(path string) (*os.File, error) {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		_ = os.MkdirAll(filepath.Dir(path), 0o777)
	}

	outfile, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open destination file %v: %w", path, err)
	}

	return outfile, nil
}

type outputBuffer struct {
	buf         *bytes.Buffer
	packageName string
	packagePath string
	// imports is a set of referenced packagePaths.
	// The value is the name from the packageMap.
	// use packageMap to retrieve the name to reference a package.
	importedPackages map[string]string
	// importPackageNames  tracks imported names. To check for collisions and do renames if necessary.
	// This is the reverse map of importedPackages.
	importPackageNames map[string]string
	packageMap         map[string]string // [<packagePath>]<packageName>
	modinfoLoader      *modinfo.Loader
}

func (b outputBuffer) ResolvePackageName(packagePath string) (string, error) {
	if packagePath == b.packagePath && b.packageName == b.packageMap[packagePath] {
		return "", nil
	}
	if packageName, found := b.importedPackages[packagePath]; found {
		return packageName, nil
	}

	packageName, found := b.packageMap[packagePath]
	if !found {
		packageMap, err := b.modinfoLoader.PackageMap([]string{packagePath})
		if err != nil {
			return "", err
		}
		var ok bool
		packageName, ok = packageMap[packagePath]
		if !ok {
			fmt.Fprintf(os.Stderr, "\tYou may have to execute: go get %v\n", packagePath)
			return "", fmt.Errorf("Could not resolve %v", packagePath)
		}
		b.packageMap[packagePath] = packageName
	}

	// packageName already taken by another import
	originalPackageName := packageName
	var i int
	_, found = b.importPackageNames[packageName]
	for found {
		i++
		packageName = fmt.Sprintf("%s%d", originalPackageName, i)
		_, found = b.importPackageNames[packageName]
	}

	b.importedPackages[packagePath] = packageName
	b.importPackageNames[packageName] = packagePath

	return packageName, nil
}

func (b outputBuffer) ResolveTypeName(packagePath string, fieldName string) (string, error) {
	pn, err := b.ResolvePackageName(packagePath)
	if err != nil {
		return "", err
	}
	if pn != "" {
		return pn + "." + fieldName, nil
	}
	return fieldName, nil
}

func usage() {
	_, _ = io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

const usageText = `wrapgen generates code for go interfaces.

Wrapgen scans all go source files for interfaces annotated with
comments: //wrapgen:generate [opts] where opts are:

        -template TEMPLATE              required. 
           Template can be a relative or absolute path.
           If there is no directoryParser separator in TEMPLATE then it's treated as name of a "builtin" template.
        -destination PATH               required.
           Path to the destination file. If multiple wrapgen:generate statements use the same
           PATH then their output is appended in the order in which they appear in the file.
        -package PACKAGE                optional.
           Go package for the target file.
           If not set then the package name is derived automatically.
        -name INTERFACE                 optional.
           Overwrites the name of the source interface.
        -vars key1=value1,k2=v2         optional.
           Injection of values into template .Vars variable.
           Usage is template specific.

wrapgen options:
`

func printVersion() {
	if bi, exists := debug.ReadBuildInfo(); exists {
		fmt.Println(bi.Main.Version)
	} else {
		slog.Error("No version information found")
	}
}
