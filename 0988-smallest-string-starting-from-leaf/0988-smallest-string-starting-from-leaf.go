/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    if root == nil {
        return ""
    }

    var minStr []byte
    helper(root, nil, &minStr)
    return string(minStr)
}

func helper(root *TreeNode, path []byte, minValue *[]byte) {
    path = append(path, byte(root.Val)+byte('a'))
    if root.Left == nil && root.Right == nil {
        currPath := reverse(path)
        currStr := string(currPath)
        
        if len(*minValue) == 0 {
            *minValue = make([]byte, len(currPath))
            copy(*minValue, currPath)
            return
        }
        
        minStr := string(*minValue)
        if minStr > currStr {
            *minValue = make([]byte, len(currPath))
            copy(*minValue, currPath)
        }
        return
    }

    if root.Left != nil {
        helper(root.Left, path, minValue)
    }
    if root.Right != nil {
        helper(root.Right, path, minValue)
    }
    
    path = path[:len(path)-1]
}

func reverse[T any](arr []T) []T {
    res := make([]T, len(arr))
    for i:=0;i<len(arr);i++{
        res[i] = arr[len(arr)-i-1]
    }
    return res
}