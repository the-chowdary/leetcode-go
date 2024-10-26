package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_numIslands(t *testing.T) {
	t.Log("Currently testing [numIslands]")
	examples := [][]string {
		{
			`[
				["1","1","1","1","0"],
				["1","1","0","1","0"],
				["1","1","0","0","0"],
				["0","0","0","0","0"]
			]`,
			`1`,
		},
		{
			`[
				["1","1","0","0","0"],
				["1","1","0","0","0"],
				["0","0","1","0","0"],
				["0","0","0","1","1"]
			]`,
			`3`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numIslands, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}