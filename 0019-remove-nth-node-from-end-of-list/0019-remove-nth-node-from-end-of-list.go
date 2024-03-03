/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // to do it in single pass:
    // - we iterate the list one by one
    // - assume we already at 3th element at c = 2 (start at 0), we want to find the 3th last
    // - then it's (c+1-th) element -> 2+1-3 = 0.
    // - then when we keep move forward, c increasing, hence the nth should also be increased
    curr := head
    var prev, nth *ListNode
    c := 0
    
    for curr != nil {
        // this is when we found the nth element the first time after c iteration
        if c + 1 - n == 0 {
            nth = head
        }else if c + 1 - n > 0 { // keep moving the nth and prev pointer
            nth, prev = nth.Next, nth
        }
        curr = curr.Next
        c++
    }

    // nth in the middle
    if prev != nil && nth != nil {
        prev.Next = nth.Next
    }else if prev == nil && nth != nil { // nth at the beginning
        head = nth.Next
    }else if prev != nil && nth == nil { // nth at the end
        prev.Next = nil
    }

    return head
}