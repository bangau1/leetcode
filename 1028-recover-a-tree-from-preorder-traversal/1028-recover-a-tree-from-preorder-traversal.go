/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(traversal string) *TreeNode {
    root, _ := dfsTraversal(traversal, 0)
    return root
}

var dash = "-"[0]
var zero, nine = "0"[0], "9"[0]
func dfsTraversal(traversal string, level int) (*TreeNode, string) {
    if len(traversal) == 0 {
        return nil, ""
    }
    if len(traversal) < level {
        return nil, traversal
    }

    for i:=0;i<level;i++{
        if traversal[i] != dash{
            return nil, traversal
        }
    }
    
    start := level
    digitStr := make([]byte, 0)
    for start < len(traversal) {
        if traversal[start] >= zero && traversal[start] <= nine {
            digitStr = append(digitStr, traversal[start])
        }else{
            break
        }
        start++
    }
    if len(digitStr) == 0 {
        return nil, traversal
    }
    
    val, _ := strconv.Atoi(string(digitStr))
    // fmt.Println(val)
    node := &TreeNode{
        Val: val,
    }
    
    left, traversal := dfsTraversal(traversal[start:], level+1)
    right, traversal := dfsTraversal(traversal, level+1)
    node.Left = left
    node.Right = right
    return node, traversal
}