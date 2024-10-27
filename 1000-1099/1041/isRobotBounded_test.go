package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_isRobotBounded(t *testing.T) {
	t.Log("Currently testing [1041 -> isRobotBounded]")
	examples := [][]string{
		{
			`"GGLLGG"`,
			`true`,
		},
		{
			`"GG"`,
			`false`,
		},
		{
			`"GL"`,
			`true`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, isRobotBounded, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
