package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_numOfWays(t *testing.T) {
	t.Log("Currently testing [1411 -> numOfWays]")
	examples := [][]string{
		{`1`, `12`},
		{`5000`, `30228214`},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numOfWays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
