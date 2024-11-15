package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_findLengthOfShortestSubarray(t *testing.T) {
	t.Log("Currently testing [1574 : findLengthOfShortestSubarray]")
	examples := [][]string{
		{
			`[1,2,3,10,4,2,3,5]`,
			`3`,
		},
		{
			`[5,4,3,2,1]`,
			`4`,
		},
		{
			`[1,2,3]`,
			`0`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, findLengthOfShortestSubarray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
