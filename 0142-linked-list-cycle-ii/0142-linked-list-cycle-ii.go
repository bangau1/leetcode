/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next 

        if fast == slow {
            // spawn another tortoise
            slow2 := head
            for slow2 != slow {
                slow2 = slow2.Next
                slow = slow.Next
            }

            return slow
        }
    }

    return nil

}