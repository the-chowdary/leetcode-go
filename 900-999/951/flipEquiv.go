/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  https://leetcode.com/problems/flip-equivalent-binary-trees

 func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }

    if root1 == nil || root2 == nil {
        return false
    }

    if root1.Val != root2.Val {
        return false
    }

    result := flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)

    if result == true {
        return true
    }

    return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
}