package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_isCircularSentence(t *testing.T) {
	t.Log("Currently testing : [2490 : isCircularSentence]")
	examples := [][]string{
		{
			`"leetcode exercises sound delightful"`,
			`true`,
		},
		{
			`"eetcode"`,
			`true`,
		},
		{
			`"Leetcode is cool"`,
			`false`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, isCircularSentence, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
