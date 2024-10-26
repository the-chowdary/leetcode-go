package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Log("currently testing [twoSum]")
	sampleIns := [][]string{{`[2,7,11,15]`, `9`}, {`[3,2,4]`, `6`}, {`[3,3]`, `6`}}
	sampleOuts := [][]string{{`[0,1]`}, {`[1,2]`}, {`[0,1]`}}
	if err := testutil.RunLeetCodeFunc(t, twoSum, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
