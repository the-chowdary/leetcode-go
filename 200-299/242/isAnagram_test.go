package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	t.Log("Currently testing [isAnagram]")
	examples := [][]string{
		{
			`"anagram"`, `"nagaram"`,
			`true`,
		},
		{
			`"rat"`, `"car"`,
			`false`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, isAnagram, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
