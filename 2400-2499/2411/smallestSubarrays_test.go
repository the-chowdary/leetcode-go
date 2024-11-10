package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_smallestSubarrays(t *testing.T) {
	t.Log("Currently testing [2411 : smallestSubarrays]")
	examples := [][]string{
		{
			`[1,0,2,1,3]`,
			`[3,3,2,2,1]`,
		},
		{
			`[1,2]`,
			`[2,1]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, smallestSubarrays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}

}
