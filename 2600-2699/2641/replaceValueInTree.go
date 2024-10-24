/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  https://leetcode.com/problems/cousins-in-binary-tree-ii
// Biweekly 102

 func replaceValueInTree(root *TreeNode) *TreeNode {
    root.Val = 0
    q := []*TreeNode{root}

    for len(q) > 0 {
        tmp := q
        q = nil

        nextLevelSum := 0
        for _, node := range tmp {
            if node.Left != nil {
                q = append(q, node.Left)
                nextLevelSum += node.Left.Val
            }

            if node.Right != nil {
                q = append(q, node.Right)
                nextLevelSum += node.Right.Val
            }
        }

        for _, node := range tmp {
            childrenSum := 0
            if node.Left != nil {
                childrenSum += node.Left.Val
            }

            if node.Right != nil {
                childrenSum += node.Right.Val
            }

            if node.Left != nil {
                node.Left.Val = nextLevelSum - childrenSum
            }

            if node.Right != nil {
                node.Right.Val = nextLevelSum - childrenSum
            }
        }
    }
    return root
    
}