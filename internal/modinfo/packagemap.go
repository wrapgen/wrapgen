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
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PackageMap returns a map of import path to package name for specified importPaths.
// For example packagePath "bla.com/client/v1" is usually named "client".
func (l *Loader) PackageMap(importPaths []string) (map[string]string, error) {
	var (
		pkgMap              = make(map[string]string, len(importPaths))
		uncachedImportPaths []string
	)

	l.packageMapLock.Lock()
	defer l.packageMapLock.Unlock()

	for _, ip := range importPaths {
		if pkgName, hit := l.packageMap[ip]; hit {
			pkgMap[ip] = pkgName
		} else {
			uncachedImportPaths = append(uncachedImportPaths, ip)
		}
	}

	if len(uncachedImportPaths) > 0 {
		// invoke go list with -e to gracefully skip packages like syscalls/js.
		args := []string{"list", "-e", "-find", "-f={{.Name}}:{{.ImportPath}}"}
		args = append(args, uncachedImportPaths...)
		cmd := exec.Command("go", args...)
		cmd.Stderr = os.Stderr
		cmd.Dir = l.moduleDir
		pipeRead, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}

		err = cmd.Start()
		if err != nil {
			return nil, fmt.Errorf(`invoking 'go list' failed: %s`, err)

		}

		s := bufio.NewScanner(pipeRead)
		for s.Scan() {
			splitted := strings.SplitN(s.Text(), ":", 2)
			if len(splitted) != 2 {
				return nil, fmt.Errorf("unexpected output from 'go list': %v", splitted)
			}
			pkgMap[splitted[1]] = splitted[0]
		}
		if err := s.Err(); err != nil {
			return nil, fmt.Errorf("reading 'go list' output resulted in error: %s", err)
		}

		err = cmd.Wait()
		if err != nil {
			return nil, fmt.Errorf(`'go list' failed: %s`, err)
		}

		for _, ip := range uncachedImportPaths {
			l.packageMap[ip] = pkgMap[ip]
		}
	}

	return pkgMap, nil
}