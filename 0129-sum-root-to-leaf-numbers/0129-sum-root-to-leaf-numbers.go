/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var sum, val int
    dfs(root, &sum, val)
    return sum
}

func dfs(root *TreeNode, sum *int, val int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *sum += val * 10 + root.Val
    }else {
        dfs(root.Left, sum, val * 10 + root.Val)
        dfs(root.Right, sum, val * 10 + root.Val)
    }
    
}