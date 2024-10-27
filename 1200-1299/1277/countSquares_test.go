package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_countSquares(t *testing.T) {
	t.Log("Currently testing [countSquares]")
	examples := [][]string{
		{
			`[
				[0,1,1,1],
				[1,1,1,1],
				[0,1,1,1]
			]`,
			`15`,
		},
		{
			`[
				[1,0,1],
				[1,1,0],
				[1,1,0]
			]`,
			`7`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, countSquares, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
