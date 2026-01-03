package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_minimumOperations(t *testing.T) {
	t.Log("Currently testing [3190 -> minimumOperations]")
	examples := [][]string{
		{
			`[1,2,3,4]`,
			`3`,
		},
		{
			`[3,6,9]`,
			`0`,
		},
		{
			`[2,4,5,7,8]`,
			`5`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimumOperations, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
