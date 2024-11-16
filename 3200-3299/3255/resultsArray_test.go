package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_resultsArray(t *testing.T) {
	t.Log("Currently testing [3254 : resultsArray]")
	examples := [][]string{
		{
			`[1,2,3,4,3,2,5]`, `3`,
			`[3,4,-1,-1,-1]`,
		},
		{
			`[2,2,2,2,2]`, `4`,
			`[-1,-1]`,
		},
		{
			`[3,2,3,2,3,2]`, `2`,
			`[-1,3,-1,3,-1]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, resultsArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
