package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	t.Log("Currently testing [isPalindrome]")
	examples := [][] string{
		{
			`"A man, a plan, a canal: Panama"`,
			`true`,
		},
		{
			`"race a car"`,
			`false`,
		},
		{
			`" "`,
			`true`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, isPalindrome, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}