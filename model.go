// Copyright 2024 Wrapgen authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"go/ast"
	"strings"
	"sync"
)

type Interface struct {
	// lock is used to ensure an interface is only read once if parsed is false.
	lock sync.Mutex
	// parsed is set to true after first call to Parse()
	parsed     bool
	fileParser *fileParser
	tps        map[string]Type
	typeSpec   *ast.TypeSpec
	Interface  *ast.InterfaceType

	Name       string
	Methods    []*Method
	AllMethods []*Method
	TypeParams []*Parameter
	Package    string
	Embedded   []*Interface
}

// ErrorInterface represent built-in error interface.
var ErrorInterface = &Interface{
	Name: "error",
	Methods: []*Method{
		{
			Name: "Error",
			Out: []*Parameter{
				{
					Name: "",
					Type: PredeclaredType("string"),
				},
			},
		},
	},
}

// Method is a single method of an interface.
type Method struct {
	Interface *Interface
	Name      string
	In, Out   []*Parameter
	Variadic  *Parameter // may be nil
}

// NamedType is an exported type in a package.
type NamedType struct {
	fileParser *fileParser
	Package    string // may be empty
	Type       string
	TypeParams *TypeParametersType
}

func (nt *NamedType) String(resolver TypeResolver) string {
	ref, err := resolver.ResolveTypeName(nt.Package, nt.Type)
	if err != nil {
		panic(err)
	}
	return ref + nt.TypeParams.String(resolver)
}

type TypeResolver interface {
	ResolveTypeName(packagePath, name string) (string, error)
}

// Parameter is an argument or return parameter of a method.
type Parameter struct {
	Name string // may be empty
	Type Type
}

// Type is a Go type.
type Type interface {
	String(resolver TypeResolver) string
}

// ArrayType is an array or slice type.
type ArrayType struct {
	Len  int // -1 for slices, >= 0 for arrays
	Type Type
}

func (at *ArrayType) String(resolver TypeResolver) string {
	s := "[]"
	if at.Len > -1 {
		s = fmt.Sprintf("[%d]", at.Len)
	}
	return s + at.Type.String(resolver)
}

// ChanType is a channel type.
type ChanType struct {
	Dir  ChanDir // 0, 1 or 2
	Type Type
}

func (ct *ChanType) String(resolver TypeResolver) string {
	s := ct.Type.String(resolver)
	if ct.Dir == RecvDir {
		return "<-chan " + s
	}
	if ct.Dir == SendDir {
		return "chan<- " + s
	}
	return "chan " + s
}

// ChanDir is a channel direction.
type ChanDir int

// Constants for channel directions.
const (
	RecvDir ChanDir = 1
	SendDir ChanDir = 2
)

// FuncType is a function type.
type FuncType struct {
	In, Out  []*Parameter
	Variadic *Parameter // may be nil
}

func (ft *FuncType) String(resolver TypeResolver) string {
	args := make([]string, len(ft.In))
	for i, p := range ft.In {
		args[i] = p.Type.String(resolver)
	}
	if ft.Variadic != nil {
		args = append(args, "..."+ft.Variadic.Type.String(resolver))
	}
	rets := make([]string, len(ft.Out))
	for i, p := range ft.Out {
		rets[i] = p.Type.String(resolver)
	}
	retString := strings.Join(rets, ", ")
	if nOut := len(ft.Out); nOut == 1 {
		retString = " " + retString
	} else if nOut > 1 {
		retString = " (" + retString + ")"
	}
	return "func(" + strings.Join(args, ", ") + ")" + retString
}

// MapType is a map type.
type MapType struct {
	Key, Value Type
}

func (mt *MapType) String(resolver TypeResolver) string {
	return "map[" + mt.Key.String(resolver) + "]" + mt.Value.String(resolver)
}

// PointerType is a pointer to another type.
type PointerType struct {
	Type Type
}

func (pt *PointerType) String(resolver TypeResolver) string {
	return "*" + pt.Type.String(resolver)
}

// PredeclaredType is a predeclared type such as "int".
type PredeclaredType string

func (pt PredeclaredType) String(_ TypeResolver) string { return string(pt) }

// TypeParametersType contains type parameters for a NamedType.
type TypeParametersType struct {
	TypeParameters []Type
}

func (tp *TypeParametersType) String(resolver TypeResolver) string {
	if tp == nil || len(tp.TypeParameters) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range tp.TypeParameters {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(v.String(resolver))
	}
	sb.WriteString("]")
	return sb.String()
}
