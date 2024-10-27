/*
Implemented treeQueries
1. Calculate height of the tree treeHeight()
2. create result array with len(height)
3. dfs the root -> update result[root.Val] = rest-height
4. dfs root.Left with max(rest_height, depth + height[root.Right])
5. dfs root.Right with max(rest_height, depth + height[root.Left])
*/
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
