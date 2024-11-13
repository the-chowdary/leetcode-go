package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_numOfSubarrays(t *testing.T) {
	t.Log("Currently testing : [1343 : numOfSubarrays]")
	examples := [][]string{
		{
			`[2,2,2,2,5,5,5,8]`, `3`, `4`,
			`3`,
		},
		{
			`[11,13,17,23,29,31,7,5,2,3]`, `3`, `5`,
			`6`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numOfSubarrays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
