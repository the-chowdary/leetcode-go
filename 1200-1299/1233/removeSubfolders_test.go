package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_removeSubfolders(t *testing.T) {
	t.Log("Currently testing [removeSubfolders]")
	examples := [][]string{
		{
			`["/a","/a/b","/c/d","/c/d/e","/c/f"]`,
			`["/a","/c/d","/c/f"]`,
		},
		{
			`["/a","/a/b/c","/a/b/d"]`,
			`["/a"]`,
		},
		{
			`["/a/b/c","/a/b/ca","/a/b/d"]`,
			`["/a/b/c","/a/b/ca","/a/b/d"]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, removeSubfolders, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
