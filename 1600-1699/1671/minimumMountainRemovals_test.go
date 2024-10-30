package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_minimumMountainRemovals(t *testing.T) {
	t.Log("Currently testing [1671 : https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array]")
	examples := [][]string{
		{
			`[1,3,1]`,
			`0`,
		},
		{
			`[2,1,1,5,6,2,3,1]`,
			`3`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimumMountainRemovals, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
