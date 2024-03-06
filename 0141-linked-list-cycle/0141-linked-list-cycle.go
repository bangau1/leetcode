/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    // can be solved by using fast-slow pointer
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        if slow == fast.Next.Next {
            return true
        }
        slow = slow.Next
        fast = fast.Next.Next
    }

    return false
}