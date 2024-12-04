package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_canMakeSubsequence(t *testing.T) {
	t.Log("Currently testing [2825 : canMakeSubsequence]")
	examples := [][]string{
		{
			`"abc"`, `"ad"`,
			`true`,
		},
		{
			`"zc"`, `"ad"`,
			`true`,
		},
		{
			`"ab"`, `"d"`,
			`false`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, canMakeSubsequence, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
