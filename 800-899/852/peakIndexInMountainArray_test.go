package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	t.Log("Currently testing [852 : peakIndexInMountainArray]")
	examples := [][]string{
		{
			`[0,1,0]`,
			`1`,
		},
		{
			`[0,2,1,0]`,
			`1`,
		},
		{
			`[0,10,5,2]`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, peakIndexInMountainArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
