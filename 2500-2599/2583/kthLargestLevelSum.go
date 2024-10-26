package main

import (
	. "leetcode-go/testutil"
	"slices"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	// BFS
	result := []int64{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		tmp := q
		q = nil
		s := int64(0)

		for _, node := range tmp {
			s += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, s)
	}
	if k > len(result) {
		return -1
	}
	slices.Sort(result)
	return result[len(result)-k]

}

func main() {

}
