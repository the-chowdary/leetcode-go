package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	t.Log("Currently testing [topKFrequent]")
	examples := [][]string {
		{
			`[1,1,1,2,2,3]`, `2`,
			`[1,2]`,
		},
		{
			`[1]`, `1`,
			`[1]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, topKFrequent, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}