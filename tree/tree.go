/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : tree.go
#   Last Modified : 2021-07-02 15:31
#   Describe      :
#
# ====================================================*/

package tree

import (
	"fmt"
	"hotbuild/strslice"
	"os"
	"path"
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
