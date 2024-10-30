package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	t.Log("Currently testing : [643 -> findMaxAverage]")
	examples := [][]string{
		{
			`[1,12,-5,-6,50,3]`, `4`,
			`12.75000`,
		},
		{
			`[5]`, `1`,
			`5.00000`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, findMaxAverage, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
