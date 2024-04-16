/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        newRoot := &TreeNode{
            Val: val,
        }
        newRoot.Left = root
        return newRoot
    }
    helper(root, val, 1, depth)
    return root
}

func helper(root *TreeNode, val int, currDepth int, depth int) *TreeNode {
    
    if currDepth + 1 == depth {
        leftNode := &TreeNode{
            Val: val,
        }
        leftNode.Left = root.Left
        root.Left = leftNode
        
        rightNode := &TreeNode{
            Val: val,
        }
        rightNode.Right = root.Right
        root.Right = rightNode
        return root
    }
    if root.Left != nil {
        helper(root.Left, val, currDepth + 1, depth)
    }
    if root.Right != nil {
        helper(root.Right, val, currDepth + 1, depth)
    }
    
    return root
}