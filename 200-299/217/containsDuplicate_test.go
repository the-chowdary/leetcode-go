package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	t.Log("Currently testing [containsDuplicate]")
	examples := [][]string {
		{
			`[1,2,3,1]`,
			`true`,
		},
		{
			`[1,2,3,4]`,
			`false`,
		},
		{
			`[1,1,1,3,3,4,3,2,4,2]`,
			`true`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, containsDuplicates, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}