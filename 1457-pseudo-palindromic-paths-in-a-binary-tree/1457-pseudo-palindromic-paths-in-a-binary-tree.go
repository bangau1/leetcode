/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pseudoPalindromicPaths (root *TreeNode) int {
    //pseudo palindromic
    // if len(path) is even, then the counter of each number is even (no odd)
    // if len(path) = n is odd, then there are only 1 number that its count is odd.
    var total int
    calculate(root, make([]int, 0), make([]int, 10), &total)
    return total
}

func calculate(root *TreeNode, path []int, counter []int, total *int) {
    path = append(path, root.Val)
    counter[root.Val]++
    if root.Left == nil && root.Right == nil{
        
        oddCount := 0
        pathLen := 0
        for i:=1;i<=9;i++{
            if counter[i] % 2 == 1 {
                oddCount++
            }
            if counter[i] > 0 {
                pathLen+=counter[i]
            }
        }
        if pathLen % 2 == 0 {
            if oddCount == 0 {
                *total+=1
            }
        }else{
            if oddCount == 1{
                *total+=1
            }
        }
       
    }else{
        if root.Left != nil {
            calculate(root.Left, path, counter, total)
        }

        if root.Right != nil {
            calculate(root.Right, path, counter, total)
        }
    }
    counter[root.Val]--
    path = path[0:len(path)-1]
}