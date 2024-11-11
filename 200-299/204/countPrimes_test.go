package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_countPrimes(t *testing.T) {
	t.Log("Currently testing [204: countPrimes]")
	examples := [][]string{
		{
			`10`,
			`4`,
		},
		{
			`0`,
			`0`,
		},
		{
			`1`,
			`0`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, countPrimes, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
