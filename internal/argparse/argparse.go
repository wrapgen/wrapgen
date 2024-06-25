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

package argparse

import (
	"fmt"
	"strings"
	"text/scanner"
)

// Split implements very basic shell-like argument splitting.
func Split(in string) ([]string, error) {
	var (
		tok     rune
		s       scanner.Scanner
		escaped bool
		quoted  bool
		buf     strings.Builder
		result  []string
	)
	s.Init(strings.NewReader(in))
	for tok != scanner.EOF {
		tok = s.Next()
		if escaped {
			switch tok {
			case '"':
				buf.WriteRune(tok)
				escaped = false
			case '\\':
				buf.WriteRune(tok)
				escaped = false
			case 'n':
				buf.WriteRune('\n')
				escaped = false
			case scanner.EOF:
				return nil, fmt.Errorf("end of string in escape sequence")
			default:
				return nil, fmt.Errorf("invalid escape sequence %q", []byte{byte(tok)})
			}
		} else {
			if quoted {
				switch tok {
				case '\\':
					escaped = true
				case '"':
					quoted = false
				case scanner.EOF:
					return nil, fmt.Errorf("end of string within quoted text")
				default:
					buf.WriteRune(tok)
				}
			} else {
				switch tok {
				case '"':
					quoted = true
				case ' ':
					result = append(result, buf.String())
					buf.Reset()
				case scanner.EOF:
					result = append(result, buf.String())
				default:
					buf.WriteRune(tok)
				}
			}
		}
	}
	return result, nil
}
