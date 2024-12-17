package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_getFinalState(t *testing.T) {
	t.Log("Currently testing [3264 : getFinalState]")
	examples := [][]string{
		{
			`[2,1,3,5,6]`, `5`, `2`,
			`[8,4,6,5,6]`,
		},
		{
			`[1,2]`, `3`, `4`,
			`[16,8]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getFinalState, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
