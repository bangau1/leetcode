/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
   var res int
   getMaxLength(root, &res)
   return res
}

func getMaxLength(root *TreeNode, maxDiameter *int) int {
     if root == nil {
        return 0
    }
    
    var left, right int
    if root.Left != nil {
        left = 1 + getMaxLength(root.Left, maxDiameter)
    }
    if root.Right != nil {
        right = 1 + getMaxLength(root.Right, maxDiameter)
    }
    *maxDiameter = max(*maxDiameter, left + right)
    return max(left, right)
}