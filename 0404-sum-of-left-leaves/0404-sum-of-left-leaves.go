/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var sum int
    if root.Left != nil {
        if root.Left.Left == nil && root.Left.Right == nil {
            sum += root.Left.Val
        }else{
            sum += sumOfLeftLeaves(root.Left)
        }
    }

    if root.Right != nil {
        sum += sumOfLeftLeaves(root.Right)
    }

    return sum
}