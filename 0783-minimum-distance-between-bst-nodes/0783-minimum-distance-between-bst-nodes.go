/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    // in bst
    // it's sorted 
    // what we need is to process it in sorted manner
    // process the left, root, then right
    var lastVal = -1
    var gDiff = math.MaxInt
    dfsTraverse(root, &lastVal, &gDiff)
    return gDiff
}

func dfsTraverse(root *TreeNode, lastVal, gDiff *int) {
    if root == nil {
        return
    }

    dfsTraverse(root.Left, lastVal, gDiff)
    if *lastVal != -1 && *gDiff > abs(root.Val - *lastVal){
        *gDiff = abs(root.Val - *lastVal)
    }
    *lastVal = root.Val
    dfsTraverse(root.Right, lastVal, gDiff)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}