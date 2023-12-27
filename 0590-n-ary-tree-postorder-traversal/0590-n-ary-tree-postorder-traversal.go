/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil {
        return nil
    }
    var res []int
    for _, child:= range root.Children{
        res = append(res, postorder(child)...)
    }
    res = append(res, root.Val)
    return res
}