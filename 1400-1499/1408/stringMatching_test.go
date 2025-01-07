package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_stringMatching(t *testing.T) {
	t.Log("Currently testing [1408 : stringMatching]")
	examples := [][]string{
		{
			`["mass","as","hero","superhero"]`,
			`["as","hero"]`,
		},
		{
			`["leetcode","et","code"]`,
			`["et","code"]`,
		},
		{
			`["blue","green","bu"]`,
			`[]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, stringMatching, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
