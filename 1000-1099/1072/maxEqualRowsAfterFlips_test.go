package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxEqualRowsAfterFilps(t *testing.T) {
	t.Log("Currently testing [1072 : maxEqualRowsAfterFlips]")
	examples := [][]string{
		{
			`[[0,1],[1,1]]`,
			`1`,
		},
		{
			`[[0,1],[1,0]]`,
			`2`,
		},
		{
			`[[0,0,0],[0,0,1],[1,1,0]]`,
			`2`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxEqualRowsAfterFlips, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
