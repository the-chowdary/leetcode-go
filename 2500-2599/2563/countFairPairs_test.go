package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_countFairPairs(t *testing.T) {
	t.Log("Currently tesing [2563 : countFairPairs]")
	examples := [][]string{
		{
			`[0,1,7,4,4,5]`, `3`, `6`,
			`6`,
		},
		{
			`[1,7,9,2,5]`, `11`, `11`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, countFairPairs1, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}

	if err := testutil.RunLeetCodeFuncWithExamples(t, countFairPairs2, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
