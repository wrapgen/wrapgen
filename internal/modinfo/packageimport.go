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

package modinfo

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

var (
	errOutsideGoPath = errors.New("source file is outside module directory")
)

// ImportPath get package import path and module-path via source file.
func ImportPath(srcDir string) (string, string, string, error) {
	srcDir, err := filepath.Abs(srcDir)
	if err != nil {
		return "", "", "", err
	}

	// trying to find the module
	currentDir := srcDir
	for {
		dat, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				// at the root
				break
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return "", "", "", err
		}
		modulePath := modfile.ModulePath(dat)
		packagePath := filepath.ToSlash(filepath.Join(modulePath, strings.TrimPrefix(srcDir, currentDir)))
		return packagePath, modulePath, currentDir, nil
	}
	return "", "", "", errOutsideGoPath
}
