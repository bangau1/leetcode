/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    // //1. recursive approach
    // if root == nil {
    //     return nil
    // }
    // var res []int
    // for _, child:= range root.Children{
    //     res = append(res, postorder(child)...)
    // }
    // res = append(res, root.Val)
    // return res

    // 2. iterative approach
    return postorderIterative(root)
}

func postorderIterative(root *Node) []int {
    if root == nil {
        return nil
    }

    var res []int
    stack := make([]*Node, 0)
    stack = append(stack, root)
    stackVal := make([]int, 0)
    for len(stack) > 0 {
        node := stack[len(stack)-1] //pop top item
        stack = stack[0:len(stack)-1]


        stackVal = append(stackVal, node.Val)

        for _, child := range node.Children{
            if child != nil {
                stack = append(stack, child)
            }
        }

    }
    for len(stackVal) > 0 {
        pop := stackVal[len(stackVal)-1]
        stackVal = stackVal[0:len(stackVal)-1]
        res = append(res, pop)
    }
    return res
}