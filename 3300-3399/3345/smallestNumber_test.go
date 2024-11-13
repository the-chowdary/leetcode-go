package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_smallestNumber(t *testing.T) {
	t.Log("Currently testing [3345 : smallestNumber]")
	examples := [][]string{
		{
			`10`, `2`,
			`10`,
		},
		{
			`15`, `3`,
			`16`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, smallestNumber, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
