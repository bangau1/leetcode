/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return root.Val
    }
    // there are 2 categories of path sum -> path that include the root or not
    // on each root, calculate both category
    a, b := helper(root)
    return max(a, b)
}

func helper(root *TreeNode) (int, int) {
    var leftSum, leftSumWithRoot, rightSum, rightSumWithRoot int
    var sumWithRoot = root.Val
    var sumWithoutRoot = math.MinInt
    if root.Left != nil {
        leftSum, leftSumWithRoot = helper(root.Left)
        sumWithRoot = max(sumWithRoot, root.Val + leftSumWithRoot)
        sumWithoutRoot = max(leftSum, leftSumWithRoot)
    }


    if root.Right != nil {
        rightSum, rightSumWithRoot = helper(root.Right)
        sumWithRoot = max(sumWithRoot, root.Val + rightSumWithRoot)
        sumWithoutRoot = max(sumWithoutRoot, rightSum, rightSumWithRoot)
    }

    if root.Left != nil && root.Right != nil {
        sumWithoutRoot = max(sumWithoutRoot, leftSumWithRoot + root.Val + rightSumWithRoot)
    }

    return sumWithoutRoot, sumWithRoot
}