package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxUniqueSplit(t *testing.T) {
	t.Log("current test executing for [MaxUniqueSplit]")
	examples := [][]string{
		{
			`"ababccc"`,
			`5`,
		},
		{
			`"aba"`,
			`2`,
		},
		{
			`"aa"`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxUniqueSplit, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
