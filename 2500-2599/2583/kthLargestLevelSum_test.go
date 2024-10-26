package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_kthLargestLevelSum(t *testing.T) {
	t.Log("Currently testing [kthLargesetLevelSum]")
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, kthLargestLevelSum, "input.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
