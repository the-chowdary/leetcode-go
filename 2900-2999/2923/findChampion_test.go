package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_findChampion(t *testing.T) {
	t.Log("Currently testing [2923 : findChampion]")
	examples := [][]string{
		{
			`[[0,1],[0,0]]`,
			`0`,
		},
		{
			`[[0,0,1],[1,0,1],[0,0,0]]`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, findChampion, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
