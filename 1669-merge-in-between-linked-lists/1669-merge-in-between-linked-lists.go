/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    var backList2 = list2
    for backList2.Next != nil {
        backList2 = backList2.Next
    }
    var aPrevNode = list1
    idx := 0
    for idx < a-1 {
        aPrevNode = aPrevNode.Next
        idx++
    }

    var bNode = aPrevNode
    for idx < b {
        bNode = bNode.Next
        idx++
    }
    aPrevNode.Next = list2
    backList2.Next = bNode.Next
    return list1


}