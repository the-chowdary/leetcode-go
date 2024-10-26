package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_treeQueries(t *testing.T) {
	t.Log("Currently testing [treeQueries]")
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, treeQueries, "input.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}