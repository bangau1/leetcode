/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    q := make([]int, 0)
    curr := head
    for curr != nil {
        q = append(q, curr.Val)
        curr = curr.Next
    }

    for len(q) > 1 {
        if q[0] != q[len(q)-1] {
            return false
        }
        q = q[1:len(q)-1]
    }
    return true
}