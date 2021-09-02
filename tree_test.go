package main

import (
	"testing"

	"github.com/wandercn/hotbuild/tree"
)

func Test_Tree(t *testing.T) {
	dir := "./"
	list := make([]string, 100)
	excludeDir := []string{"vendor", ".git", "docs"}
	err := tree.TreeDirs(dir, &list, excludeDir)
	if err != nil {
		t.Errorf("test tree failed: %v", err)
	}
}
