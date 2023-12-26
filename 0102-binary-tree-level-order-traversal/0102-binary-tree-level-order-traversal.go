/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    writeLevelOrder(root, 1, &result)

    return result
}

func writeLevelOrder(root *TreeNode, level int, result *[][]int) {
    if root == nil {
        return
    }
    if len(*result) < level {
        *result = append(*result, make([]int, 0))
    }

    (*result)[level-1] = append((*result)[level-1], root.Val)

    writeLevelOrder(root.Left, level + 1, result)
    writeLevelOrder(root.Right, level + 1, result)
}