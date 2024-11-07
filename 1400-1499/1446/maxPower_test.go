package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxPower(t *testing.T) {
	t.Log("Currently testing [1446 : maxPower]")
	examples := [][]string{
		{
			`"leetcode"`,
			`2`,
		},
		{
			`"abbcccddddeeeeedcba"`,
			`5`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxPower, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
