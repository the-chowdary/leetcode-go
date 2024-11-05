package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_minChanges(t *testing.T) {
	t.Log("Currently testing : [2914 : minChanges]")
	exapmles := [][]string{
		{
			`"1001"`,
			`2`,
		},
		{
			`"10"`,
			`1`,
		},
		{
			`"0000"`,
			`0`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minChanges, exapmles, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
