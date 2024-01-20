/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    return rangeSum(root, low, high, 0)
}

func rangeSum(root *TreeNode, low, high int, currentSum int) int{
    if root == nil {
        return currentSum
    }
    if root.Val < low {
        return rangeSum(root.Right, low, high, currentSum)
    }else if root.Val > high {
        return rangeSum(root.Left, low, high, currentSum)
    }else{
        currentSum += root.Val
        currentSum += rangeSum(root.Left, low, high, 0)
        currentSum += rangeSum(root.Right, low, high, 0)
    }

    return currentSum
}