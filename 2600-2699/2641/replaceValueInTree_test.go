package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_replaceValueInTree(t *testing.T) {
	t.Log("Currently testing [replaceValueInTree]")
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, replaceValueInTree, "input.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}