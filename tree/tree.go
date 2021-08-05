/*
   Copyright (c) 2021 ffactory.org
   hotbuild is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/
package tree

import (
	"fmt"
	"os"
	"path"

	"github.com/wandercn/hotbuild/strslice"
)

func TreeDirs(dir string, dirList *[]string, excludeDir []string) error {
	d, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read dir failed: %v", err)
	}

	for _, sub := range d {
		if sub.IsDir() && !strslice.IsStrInSlice(sub.Name(), excludeDir) {
			*dirList = append(*dirList, path.Join(dir, sub.Name()))
			TreeDirs(path.Join(dir, sub.Name()), dirList, excludeDir)
		}
	}
	return nil
}

func IsExistDir(dir string) bool {
	d, err := os.Stat(dir)
	if err != nil {
		return false
	}
	if d.IsDir() {
		return true
	} else {
		return false
	}
}
