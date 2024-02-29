/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    latestVal := make(map[int]int)
    bfs := make([]node, 0)
    if root != nil {
        bfs = append(bfs, node{root.Val, 0, root})
    }

    for len(bfs) > 0 {
        cur := bfs[0]
        bfs = bfs[1:]

        if cur.level % 2 == 0 {
            if !putEvenLevel(cur, latestVal){
                return false
            }
        }else{
            if !putOddLevel(cur, latestVal){
                return false
            }
        }
        if cur.n.Left != nil {
            bfs = append(bfs, node{cur.n.Left.Val, cur.level + 1, cur.n.Left})
        }
        if cur.n.Right != nil {
            bfs = append(bfs, node{cur.n.Right.Val, cur.level + 1, cur.n.Right})
        }
    }

    return true
    
}

func putEvenLevel(node node, latestVal map[int]int) bool {
    if node.val % 2 == 0 {
        return false
    }
    if _, ok:=latestVal[node.level]; !ok {
        latestVal[node.level] = node.val
        return true 
    }

    if latestVal[node.level] >= node.val {
        return false
    }
    latestVal[node.level] = node.val
    return true
}

func putOddLevel(node node, latestVal map[int]int) bool {
    if node.val % 2 == 1 {
        return false
    }
    if _, ok:=latestVal[node.level]; !ok {
        latestVal[node.level] = node.val
        return true 
    }

    if latestVal[node.level] <= node.val {
        return false
    }
    latestVal[node.level] = node.val
    return true
}

type node struct {
    val, level int
    n *TreeNode
}