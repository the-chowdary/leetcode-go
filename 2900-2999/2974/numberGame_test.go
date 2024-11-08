package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_numberGame(t *testing.T) {
	t.Log("Currently testing [2974: numberGame]")
	examples := [][]string{
		{
			`[5,4,2,3]`,
			`[3,2,5,4]`,
		},
		{
			`[2,5]`,
			`[5,2]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numberGame, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
