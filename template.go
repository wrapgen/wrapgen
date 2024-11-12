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
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/wrapgen/wrapgen/internal/cache"
)

var (
	//go:embed template
	builtinTemplates embed.FS
	templateCache    = cache.New[string, *template.Template]()
)

func getTemplate(templateBasedir string, name string) (*template.Template, error) {
	var cacheKey string
	if strings.Contains(name, "/") {
		cacheKey = filepath.Join(templateBasedir, name)
	} else {
		cacheKey = name
	}

	return templateCache.GetOrAdd(cacheKey, func(cacheKey string) (*template.Template, error) {
		t := template.New(cacheKey)

		var templateText []byte
		var err error
		if strings.Contains(name, string([]byte{filepath.Separator})) {
			templateText, err = os.ReadFile(cacheKey)
			if err != nil {
				return nil, fmt.Errorf("failed to open template file: %s", err)
			}
		} else {
			templateText, err = builtinTemplates.ReadFile(filepath.Join("template", name))
			if err != nil {
				return nil, fmt.Errorf("failed to open builtin template: %s", err)
			}
		}

		t.Funcs(templateFunctions(nil))
		_, err = t.Parse(string(templateText))
		if err != nil {
			return nil, fmt.Errorf("failed to parse template: %s", err)
		}

		return t, nil
	})
}

func templateFunctions(output *outputBuffer) map[string]any {
	return map[string]any{
		`formatType`: func(p any) (string, error) {
			switch v := p.(type) {
			case Type:
				return v.String(output), nil
			case []*Parameter:
				b := strings.Builder{}
				if len(v) > 0 {
					b.WriteString("[")
					for i, tp := range v {
						b.WriteString(tp.Name)
						b.WriteString(" ")
						b.WriteString(tp.Type.String(output))
						if i < len(v)-1 {
							b.WriteString(", ")
						}
					}
					b.WriteString("]")
				}
				return b.String(), nil
			case *Interface:
				b := strings.Builder{}
				ref, err := output.ResolveTypeName(v.Package, v.Name)
				if err != nil {
					return "", err
				}
				b.WriteString(ref)
				if len(v.TypeParams) > 0 {
					b.WriteString("[")
					for i, tp := range v.TypeParams {
						b.WriteString(tp.Name)
						if i < len(v.TypeParams)-1 {
							b.WriteString(", ")
						}
					}
					b.WriteString("]")
				}
				return b.String(), nil
			default:
				return "", fmt.Errorf("unsupported type %T for formatType", p)
			}
		},
		`import`: func(packagePath string) (string, error) {
			packageName, err := output.ResolvePackageName(packagePath)
			return packageName, err
		},
		`parameterNamesPlain`: func(prefix string, p []*Parameter) []string {
			e := make([]string, 0, len(p))
			for i, param := range p {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				e = append(e, name)
			}
			return e
		},
		`parametersOutPlain`: func(prefix string, m *Method) string {
			e := make([]string, 0, len(m.Out))
			for i, param := range m.Out {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				e = append(e, fmt.Sprintf("%s %s", name, param.Type.String(output)))
			}
			return strings.Join(e, ", ")
		},
		`parameterNames`: func(prefix string, p []*Parameter) []string {
			e := make([]string, 0, len(p))
			for i, param := range p {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				if i == 0 {
					if nt, ok := param.Type.(*NamedType); ok && nt.Package == "context" && nt.Type == "Context" {
						name = "ctx"
					}
				}
				if i == len(p)-1 {
					if predeclT, ok := param.Type.(PredeclaredType); ok && string(predeclT) == "error" {
						name = "err"
					}
				}
				e = append(e, name)
			}
			return e
		},
		`parameterNameVariadic`: func(prefix string, m *Method) string {
			if m.Variadic == nil {
				return ""
			}
			name := m.Variadic.Name
			if name == "" {
				name = fmt.Sprintf("%s%d", prefix, len(m.In))
			}
			return name
		},
		`parametersInNames`: func(prefix string, m *Method) []string {
			e := make([]string, 0)
			for i, param := range m.In {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				if i == 0 {
					if nt, ok := param.Type.(*NamedType); ok && nt.Package == "context" && nt.Type == "Context" {
						name = "ctx"
					}
				}
				if i == len(m.In)-1 {
					if predeclT, ok := param.Type.(PredeclaredType); ok && string(predeclT) == "error" {
						name = "err"
					}
				}
				e = append(e, name)
			}
			if m.Variadic != nil {
				name := m.Variadic.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, len(m.In))
				}
				e = append(e, name)
			}
			return e
		},
		`parametersInNamesUnpackVariadic`: func(prefix string, m *Method) string {
			e := make([]string, 0)
			for i, param := range m.In {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				if i == 0 {
					if nt, ok := param.Type.(*NamedType); ok && nt.Package == "context" && nt.Type == "Context" {
						name = "ctx"
					}
				}
				if i == len(m.In)-1 {
					if predeclT, ok := param.Type.(PredeclaredType); ok && string(predeclT) == "error" {
						name = "err"
					}
				}
				e = append(e, name)
			}
			if m.Variadic != nil {
				name := m.Variadic.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, len(m.In))
				}
				e = append(e, fmt.Sprintf("%s...", name))
			}
			return strings.Join(e, ", ")
		},
		`parametersIn`: func(prefix string, m *Method) string {
			e := make([]string, 0)
			for i, param := range m.In {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				if i == 0 {
					if nt, ok := param.Type.(*NamedType); ok && nt.Package == "context" && nt.Type == "Context" {
						name = "ctx"
					}
				}
				if i == len(m.In)-1 {
					if predeclT, ok := param.Type.(PredeclaredType); ok && string(predeclT) == "error" {
						name = "err"
					}
				}
				e = append(e, fmt.Sprintf("%s %s", name, param.Type.String(output)))
			}
			if m.Variadic != nil {
				name := m.Variadic.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, len(m.In))
				}
				e = append(e, fmt.Sprintf("%s ...%s", name, m.Variadic.Type.String(output)))
			}
			return strings.Join(e, ", ")
		},
		`parametersInAnyType`: func(prefix string, m *Method) string {
			e := make([]string, 0)
			for i, param := range m.In {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				if i == 0 {
					if nt, ok := param.Type.(*NamedType); ok && nt.Package == "context" && nt.Type == "Context" {
						name = "ctx"
					}
				}
				if i == len(m.In)-1 {
					if predeclT, ok := param.Type.(PredeclaredType); ok && string(predeclT) == "error" {
						name = "err"
					}
				}
				if i == len(m.In)-1 {
					e = append(e, fmt.Sprintf("%s any", name))
				} else {
					e = append(e, name)
				}
			}
			if m.Variadic != nil {
				name := m.Variadic.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, len(m.In))
				}
				e = append(e, fmt.Sprintf("%s ...any", name))
			}
			return strings.Join(e, ", ")
		},
		`parametersOut`: func(prefix string, m *Method) string {
			e := make([]string, 0, len(m.Out))
			for i, param := range m.Out {
				name := param.Name
				if name == "" {
					name = fmt.Sprintf("%s%d", prefix, i)
				}
				if i == 0 {
					if nt, ok := param.Type.(*NamedType); ok && nt.Package == "context" && nt.Type == "Context" {
						name = "ctx"
					}
				}
				if i == len(m.Out)-1 {
					if predeclT, ok := param.Type.(PredeclaredType); ok && string(predeclT) == "error" {
						name = "err"
					}
				}
				e = append(e, fmt.Sprintf("%s %s", name, param.Type.String(output)))
			}
			return strings.Join(e, ", ")
		},
		`parametersTypesOnly`: func(p []*Parameter) string {
			e := make([]string, 0, len(p))
			for _, param := range p {
				e = append(e, param.Type.String(output))
			}
			return strings.Join(e, ", ")
		},
		`parametersInTypes`: func(m *Method) string {
			var e []string
			for _, param := range m.In {
				e = append(e, param.Type.String(output))
			}
			if m.Variadic != nil {
				e = append(e, "..."+m.Variadic.Type.String(output))
			}
			return strings.Join(e, ", ")
		},
		`parametersOutTypes`: func(m *Method) string {
			if len(m.Out) == 0 {
				return ""
			}
			var e []string
			for _, param := range m.Out {
				e = append(e, param.Type.String(output))
			}
			out := strings.Join(e, ", ")
			if len(e) > 1 {
				return " (" + out + ")"
			}
			return " " + out
		},
		`uniqueIdentifier`: func(base string, m *Method) string {
			var others []string
			for _, in := range m.In {
				others = append(others, in.Name)
			}
			for _, out := range m.Out {
				others = append(others, out.Name)
			}
			if m.Variadic != nil {
				others = append(others, m.Variadic.Name)
			}
			number := 2
			newName := base
			for slices.Contains(others, newName) {
				newName = fmt.Sprintf("%s_%d", base, number)
				number++
			}
			return newName
		},
		`typeParameters`: func(intf *Interface) string {
			if len(intf.TypeParams) == 0 {
				return ``
			}
			var t []string
			for _, tp := range intf.TypeParams {
				t = append(t, fmt.Sprintf("%s %s", tp.Name, tp.Type.String(output)))
			}
			return "[" + strings.Join(t, ",") + "]"
		},
		`typeParameterNames`: func(intf *Interface) string {
			if len(intf.TypeParams) == 0 {
				return ``
			}
			var t []string
			for _, tp := range intf.TypeParams {
				t = append(t, tp.Name)
			}
			return "[" + strings.Join(t, ",") + "]"
		},
		`hasKey`: func(d map[string]interface{}, key string) bool {
			_, ok := d[key]
			return ok
		},
		`list`: func(v ...string) []string {
			return v
		},
		`join`: func(sep string, v []string) string {
			return strings.Join(v, sep)
		},
		`push`: func(list []string, v string) []string {
			out := make([]string, 0, len(list)+1)
			out = append(out, list...)
			out = append(out, v)
			return out
		},
		`concat`: func(lists ...[]string) []string {
			var out []string
			for _, l := range lists {
				out = append(out, l...)
			}
			return out
		},
		`last`: func(l []string) string {
			if len(l) == 0 {
				panic("get last element from empty list")
			}
			return l[len(l)-1]
		},
		`title`: strings.Title, // nolint:staticcheck
	}
}
