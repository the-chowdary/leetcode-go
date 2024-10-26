package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	t.Log("Currently testing [groupAnagrams]")
	examples := [][]string {
		// {
		// 	`["eat","tea","tan","ate","nat","bat"]`,
		// 	`[["bat"],["nat","tan"],["ate","eat","tea"]]`,
		// }, 
		{
			`[""]`,
			`[[""]]`,
		},
		{
			`["a"]`,
			`[["a"]]`,
		},
	}
	// sampleIns := [][]string {{``}, {`[""]`}, {`["a"]`}}
	// sampleOuts := [][]string {{`[["bat"],["ate","eat","tea"],["nat","tan"]]`}, {`[[""]]`}, {`[["a"]]`}}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, groupAnagrams, examples, targetCaseNum); err != nil {
		t.Fatal("err")
	}
}