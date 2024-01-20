/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaf1 := getLeaf(root1)
    leaf2 := getLeaf(root2)
    if len(leaf1) != len(leaf2){
        return false
    }
    for i:=0;i<len(leaf1);i++{
        if leaf1[i] != leaf2[i] {
            return false
        }
    }
    return true
}

func getLeaf(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }

    var res []int
    res = append(res, getLeaf(root.Left)...)
    res = append(res, getLeaf(root.Right)...)
    return res
}