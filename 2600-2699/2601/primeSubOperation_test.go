package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_primeSubOperation(t *testing.T) {
	t.Log("Currently testing [2601: primeSubOperation]")
	examples := [][]string{
		{
			`[4,9,6,10]`,
			`true`,
		},
		{
			`[6,8,11,12]`,
			`true`,
		},
		{
			`[5,8,3]`,
			`false`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, primeSubOperation, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
