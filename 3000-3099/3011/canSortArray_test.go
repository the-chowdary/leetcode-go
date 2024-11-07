package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_canSortArray(t *testing.T) {
	t.Log("Currently testing [3011 : canSortArray]")
	examples := [][]string{
		{
			`[8,4,2,30,15]`,
			`true`,
		},
		{
			`[1,2,3,4,5]`,
			`true`,
		},
		{
			`[3,16,8,4,2]`,
			`false`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, canSortArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
