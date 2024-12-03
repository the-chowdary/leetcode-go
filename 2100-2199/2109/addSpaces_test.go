package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_addSpaces(t *testing.T) {
	t.Log("Currently testing [2109 : addSpaces]")
	examples := [][]string{
		{
			`"LeetcodeHelpsMeLearn"`, `[8,13,15]`,
			`"Leetcode Helps Me Learn"`,
		},
		{
			`"icodeinpython"`, `[1,5,7,9]`,
			`"i code in py thon"`,
		},
		{
			`"spacing"`, `[0,1,2,3,4,5,6]`,
			`" s p a c i n g"`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, addSpaces, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
