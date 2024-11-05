package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_rotateString(t *testing.T) {
	t.Log("Currently testing [796 : rotateString] ")
	examples := [][]string{
		{
			`"abcde"`, `"cdeab"`,
			`true`,
		},
		{
			`"abcde"`, `"abced"`,
			`false`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, rotateString, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
