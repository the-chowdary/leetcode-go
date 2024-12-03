package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_minSetSize(t *testing.T) {
	t.Log("Currently testing [1338 : minSetSize]")
	examples := [][]string{
		{
			`[3,3,3,3,5,5,5,2,2,7]`,
			`2`,
		},
		{
			`[7,7,7,7,7,7]`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minSetSize, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
