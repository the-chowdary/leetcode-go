package main

import . "leetcode-go/testutil"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func treeQueries(root *TreeNode, queries []int) []int {
	height := map[*TreeNode]int{}

	var treeHeight func(root *TreeNode) int

	treeHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		height[root] = 1 + max(treeHeight(root.Left), treeHeight(root.Right))
		return height[root]
	}
	treeHeight(root)

	result := make([]int, len(height)+1)

	var dfs func(root *TreeNode, depth int, restHeight int)

	dfs = func(root *TreeNode, depth int, restHeight int) {
		if root == nil {
			return
		}
		depth++
		result[root.Val] = restHeight
		dfs(root.Left, depth, max(restHeight, depth+height[root.Right]))
		dfs(root.Right, depth, max(restHeight, depth+height[root.Left]))
	}
	dfs(root, -1, 0)

	for i, q := range queries {
		queries[i] = result[q]
	}
	return queries
}

func main() {

}
