/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func frequenciesOfElements(head *ListNode) *ListNode {
    counter := make(map[int]*ListNode)
    dummy := &ListNode {}
    curr := head

    var item, prev *ListNode
    var found bool

    prev = dummy
    for curr != nil {
        item, found = counter[curr.Val]
        if found {
            item.Val+=1
        } else{
            prev.Next = &ListNode{
                Val: 1,
            }
            counter[curr.Val] = prev.Next
            prev = prev.Next
        }
        curr = curr.Next
    }

    return dummy.Next
}