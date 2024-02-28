/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    var res = data{-1,-1}
    dfs(root, 0, &res)
    return res.val
}

func dfs(node *TreeNode, level int, result *data) {
    if node == nil {
        return
    }

    if node.Left == nil && node.Right == nil {
        if level > result.level {
            result.val = node.Val
            result.level = level
        }       
        return
    }
    dfs(node.Left, level + 1, result)
    dfs(node.Right, level + 1, result)
}

type data struct {
    val int
    level int
}