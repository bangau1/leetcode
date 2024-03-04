/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func gameResult(head *ListNode) string {
    scores := [2]int{} // 0 -> even team, 1 -> odd team

    curr := head
    for curr != nil {
        if curr.Val > curr.Next.Val {
            scores[0]++
        }else if curr.Val < curr.Next.Val {
            scores[1]++
        } 
        curr = curr.Next.Next
    }
    if scores[0] > scores[1] {
        return "Even"
    }else if scores[0] < scores[1]{
        return "Odd"
    }
    return "Tie"

}