package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_sumFourDivisors(t *testing.T) {
	t.Log("Currently testing [1390 : sumFourDivisors]")
	examples := [][]string{
		{
			`[21,4,7]`,
			`32`,
		},
		{
			`[1,2,3,4,5]`,
			`0`,
		},
		{
			`[21, 21]`,
			`64`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, sumFourDivisors, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
