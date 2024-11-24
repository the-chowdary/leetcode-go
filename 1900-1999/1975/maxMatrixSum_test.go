package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxMatrixSum(t *testing.T) {
	t.Log("Currently testing [1975 : maxMatrixSum]")
	examples := [][]string{
		{
			`[[1,-1],[-1,1]]`,
			`4`,
		},
		{
			`[[1,2,3],[-1,-2,-3],[1,2,3]]`,
			`16`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxMatrixSum, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
