package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_closestPrimes(t *testing.T) {
	t.Log("Currently testing [2523 : closestPrimes]")
	examples := [][]string{
		{
			`10`, `19`,
			`[11,13]`,
		},
		{
			`4`, `-6`,
			`[-1,-1]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, closestPrimes, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
