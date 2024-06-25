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
	"slices"
	"testing"
)

func TestSplit(t *testing.T) {
	for input, expectedOutput := range map[string][]string{
		"foo bar baz":       {"foo", "bar", "baz"},
		`unqouted "qouted"`: {`unqouted`, `qouted`},
		`unqouted\\`:        {`unqouted\\`},
		`"qouted\\"`:        {`qouted\`},
		`"qouted\""`:        {`qouted"`},
		`unqouted \a`:       {`unqouted`, `\a`},
		`unqouted \n`:       {`unqouted`, `\n`},
		`qouted "\n"`:       {`qouted`, "\n"},
	} {
		actualOutput, err := Split(input)
		if err != nil {
			t.Fatalf("input %q split failed: %s", input, err)
		}
		if !slices.Equal(expectedOutput, actualOutput) {
			t.Fatalf("expected: %q, actual: %q", expectedOutput, actualOutput)
		}
	}

	for input, expectedError := range map[string]string{
		`foo "bar`:          "end of string within quoted text",
		`foo "\a"`:          `invalid escape sequence "a"`,
		`"unexpected end \`: `end of string in escape sequence`,
	} {
		_, err := Split(input)
		if err == nil {
			t.Fatalf("expected error on input: %v", input)
		}
		if err.Error() != expectedError {
			t.Fatalf("expected: %s, actual: %s", expectedError, err.Error())
		}
	}
}
