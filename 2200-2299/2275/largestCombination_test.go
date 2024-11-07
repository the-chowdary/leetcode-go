package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_largestCombination(t *testing.T) {
	t.Log("Currently testing [2275 : largestCombination]")
	examples := [][]string{
		{
			`[16,17,71,62,12,24,14]`,
			`4`,
		},
		{
			`[8,8]`,
			`2`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, largestCombination, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
