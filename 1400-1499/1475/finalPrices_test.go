package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_finalPrices(t *testing.T) {
	t.Log("Currently testing [1475 : finalPrices]")
	examples := [][]string{
		{
			`[8,4,6,2,3]`,
			`[4,2,4,2,3]`,
		},
		{
			`[1,2,3,4,5]`,
			`[1,2,3,4,5]`,
		},
		{
			`[10,1,1,6]`,
			`[9,0,1,6]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, finalPrices, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
