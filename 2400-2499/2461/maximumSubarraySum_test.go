package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maximumSubarraySum(t *testing.T) {
	t.Log("Currently testing [2461 : maximumSubarraySum]")
	examples := [][]string{
		{
			`[1,5,4,2,9,9,9]`, `3`,
			`15`,
		},
		{
			`[4,4,4]`, `3`,
			`0`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maximumSubarraySum, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
