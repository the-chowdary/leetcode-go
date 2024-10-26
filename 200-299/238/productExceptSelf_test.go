package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	t.Log("Currently testing [productExceptSelf]")
	examples := [][]string{
		{
			`[1,2,3,4]`,
			`[24,12,8,6]`,
		},
		{
			`[-1,1,0,-3,3]`,
			`[0,0,9,0,0]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, productExceptSelf, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
