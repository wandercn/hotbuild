package main

import (
	"fmt"
	"hotbuild/tree"
	"testing"
)

func Test_Tree(t *testing.T) {

	dir := "/Users/apple/workspace/src/"
	list := make([]string, 100)
	excludeDir := []string{"vendor", ".git", "docs"}
	err := tree.TreeDirs(dir, &list, excludeDir)
	if err != nil {
		t.Fatalf("test tree failed: %v", err)
	} else {
		for _, d := range list {
			fmt.Println(d)
			t.Log(d)
		}
	}

}
