package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	t.Log("Currently testing [longestConsecutive]")
	examples := [][]string {
		{
			`[100,4,200,1,3,2]`,
			`4`,
		},
		{
			`[0,3,7,2,5,8,4,6,0,1]`,
			`9`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, longestConsecutive, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}