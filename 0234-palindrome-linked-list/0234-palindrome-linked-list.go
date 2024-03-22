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
    l, r := 0, len(q)-1
    for l < r {
        if q[l] != q[r]{
            return false
        }
        l++
        r--
    }
    return true
}