package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_isAcronym(t *testing.T) {
	t.Log("Currently testing [2828 : isAcronym]")
	examples := [][]string{
		{
			`["alice","bob","charlie"]`, `"abc"`,
			`true`,
		},
		{
			`["an","apple"]`, `"a"`,
			`false`,
		},
		{
			`["never","gonna","give","up","on","you"]`, `"ngguoy"`,
			`true`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, isAcronym, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
