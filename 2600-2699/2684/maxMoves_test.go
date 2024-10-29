package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxMoves(t *testing.T) {
	t.Log("Currently testing [2648 : maxMoves]")

	examples := [][]string{
		{
			`[[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]`,
			`3`,
		},
		{
			`[[3,2,4],[2,1,9],[1,1,7]]`,
			`0`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxMoves, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
