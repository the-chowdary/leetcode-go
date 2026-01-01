package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_PlusOne(t *testing.T) {
	examples := [][]string{
		{
			`[1,2,3]`,
			`[1,2,4]`,
		},
		{
			`[4,3,2,1]`,
			`[4,3,2,2]`,
		},
		{
			`[9]`,
			`[1,0]`,
		},
	}

	t.Run("Currently testing [66 -> plusOne]", func(t *testing.T) {
		targetCaseNum := 0
		if err := testutil.RunLeetCodeFuncWithExamples(t, plusOne, examples, targetCaseNum); err != nil {
			t.Fatal(err)
		}
	})
}
