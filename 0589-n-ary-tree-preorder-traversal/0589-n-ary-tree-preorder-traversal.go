/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type note struct{
    node *Node
    childIdx int
}

func preorder(root *Node) []int {
    if root == nil {
        return nil
    }

    // iterative approach
    var res []int
    var stack []*note
    stack = append(stack, &note{node: root, childIdx:0})
    for len(stack) > 0 {
        u := stack[len(stack)-1]

        if len(u.node.Children) == 0 || u.childIdx == 0 {
            res = append(res, u.node.Val)
        }

        if u.childIdx < len(u.node.Children) {
            newNote := &note{node: u.node.Children[u.childIdx], childIdx: 0}
            stack = append(stack, newNote)
            u.childIdx++
        }else{
            stack = stack[0:len(stack)-1]
        }
    }
   
    return res
}