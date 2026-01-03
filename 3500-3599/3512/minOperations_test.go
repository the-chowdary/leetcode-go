package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_minOperations(t *testing.T) {
	t.Log("Currently testing [3512 -> minOperations]")
	examples := [][]string{
		{
			`[3,9,7]`, `5`,
			`4`,
		},
		{
			`[4,1,3]`, `4`,
			`0`,
		},
		{
			`[3,2]`, `6`,
			`5`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minOperations, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
