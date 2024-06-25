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
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

var flagOverwrite = flag.Bool("overwrite", false, "Overwrite expected files")

func mapWriter() (func(path string) (stringWriteCloser, error), map[string][]byte) {
	var (
		l sync.Mutex
		m = map[string][]byte{}
	)

	return func(path string) (stringWriteCloser, error) {
		l.Lock()
		defer l.Unlock()

		_, exists := m[path]
		if exists {
			return nil, fmt.Errorf("file %s already exists", path)
		}
		return &mapWriterWriter{
			l:    &l,
			m:    m,
			path: path,
		}, nil
	}, m
}

type mapWriterWriter struct {
	l    *sync.Mutex
	m    map[string][]byte
	path string
}

func (m mapWriterWriter) Write(p []byte) (n int, err error) {
	m.l.Lock()
	defer m.l.Unlock()
	m.m[m.path] = append(m.m[m.path], p...)
	return len(p), nil
}

func (m mapWriterWriter) Close() error {
	return nil
}

func (m mapWriterWriter) WriteString(s string) (n int, err error) {
	return m.Write([]byte(s))
}

func TestExamples(t *testing.T) {
	w, m := mapWriter()
	processDirectory("testdata/examples/", w)

	for writtenPath, writtenFileBytes := range m {
		t.Run(strings.TrimPrefix(writtenPath, "testdata/examples/"), func(t *testing.T) {
			expectedBytes, err := os.ReadFile(writtenPath)
			if err != nil {
				t.Fatalf("written to unexpected file %v: %s", writtenPath, err)
			}

			if diff := diff(writtenPath, string(expectedBytes), string(writtenFileBytes)); diff != "" {
				if *flagOverwrite {
					_ = os.WriteFile(writtenPath, writtenFileBytes, 0666)
					t.Fatalf("Overwriting %v", writtenPath)
				}
				t.Fatalf("File content differ %s:\n%s", writtenPath, diff)
			}
			_ = writtenFileBytes
		})
	}
}

// diff is a totally simple text-diff function that returns empty-string if text1==text2.
func diff(name, expected, actual string) string {
	expectedLines, actualLines := strings.Split(expected, "\n"), strings.Split(actual, "\n")
	for i := 0; i < len(expectedLines) && i < len(actualLines); i++ {
		if expectedLines[i] == actualLines[i] {
			continue
		}
		return fmt.Sprintf(`Error: Not equal: %s:%d
       	            	expected: %q
       	            	actual  : %q`, name, i+1, expectedLines[i], actualLines[i])
	}
	if len(expectedLines) > len(actualLines) {
		return fmt.Sprintf(`Error: Not equal: %s:%d
       	            	expected: %q
       	            	actual  : EOF`, name, len(expectedLines), expectedLines[len(expectedLines)-1])
	}
	if len(expectedLines) < len(actualLines) {
		return fmt.Sprintf(`Error: Not equal: %s:%d
       	            	expected: EOF
       	            	actual  : %s`, name, len(expectedLines), actualLines[len(actualLines)-1])
	}

	return ``
}
