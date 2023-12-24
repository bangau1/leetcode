type minHeap []*ListNode
func (m minHeap) Len() int {
  return len(m)
}

func (m minHeap) Less(a, b int) bool {
  return m[a].Val < m[b].Val
}
func (m minHeap) Swap(a, b int) {
  m[a], m[b] = m[b], m[a]
}

func (m *minHeap) Pop() any {
  l := m.Len()
  item := (*m)[l-1]
  *m = (*m)[:l-1]
  return item
}

func (m *minHeap) Push(a any){
  *m = append(*m, a.(*ListNode))
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    heads := make(minHeap, 0)

    for _, l := range lists {
        if l != nil {
        heads = append(heads, l)
        }
    }
    heap.Init(&heads)

    
    var result, currHead, prevHead *ListNode

    for heads.Len() > 0 {
        prevHead = currHead
        // we got the first element from the heads (the smallest from the heads list)
        currHead =  heap.Pop(&heads).(*ListNode)

        if prevHead == nil {
            result = currHead // we appoint the result the first head/smallest item
        }else{
            prevHead.Next = currHead
        }

        // push the next item from currentHead to the minHeap
        if currHead.Next != nil {
        heap.Push(&heads, currHead.Next)
        }
    }

    return result
}
