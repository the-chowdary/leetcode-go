package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_repeatedNTimes(t *testing.T) {
	examples := [][]string{
		{
			`[1,2,3,3]`,
			`3`,
		},
		{
			`[2,1,2,5,3,2]`,
			`2`,
		},
		{
			`[5,1,5,2,5,3,5,4]`,
			`5`,
		},
	}
	t.Run("Currently testing [961 -> repeatedNTimes]", func(t *testing.T) {
		targetCaseNum := 0
		if err := testutil.RunLeetCodeFuncWithExamples(t, repeatedNTimes, examples, targetCaseNum); err != nil {
			t.Fatal(err)
		}
	})
}
