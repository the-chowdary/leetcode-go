package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_flipEquiv(t *testing.T) {
	t.Log("Currently testing [flipEquiv]")
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, flipEquiv, "input.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
