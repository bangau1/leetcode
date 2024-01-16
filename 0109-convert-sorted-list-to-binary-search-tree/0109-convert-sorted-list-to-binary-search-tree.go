/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type AvlTree struct {
    root *AvlNode
}

func (t *AvlTree) Insert(val int) {
    t.root = t.insert(t.root, val)
}

func (t *AvlTree) LeftRotate(node *AvlNode) *AvlNode {
    // A
    //    -
    //      B
    //    -
    //  Y

    A := node
    B := node.Right
    Y := B.Left

    B.Left = A
    A.Right = Y


    A.SetHeight(1 + max(A.Left.GetHeight(), A.Right.GetHeight()))
    B.SetHeight(1 + max(B.Left.GetHeight(), B.Right.GetHeight()))
    return B
}

func (t *AvlTree) RightRotate(node *AvlNode) *AvlNode {
    //      B
    //    -
    //  A
    //    -
    //      Y
    A := node.Left
    B := node
    Y := A.Right
 
    A.Right = B
    B.Left = Y

    A.SetHeight(1 + max(A.Left.GetHeight(), A.Right.GetHeight()))
    B.SetHeight(1 + max(B.Left.GetHeight(), B.Right.GetHeight()))

    return A
}

func (t *AvlTree) insert(root *AvlNode, val int) *AvlNode {
    // base case
    if root == nil {
        return NewAvlNode(val)
    }

    if val < root.Val {
        root.Left = t.insert(root.Left, val)
    }else {
        root.Right = t.insert(root.Right, val)
    }

    root.SetHeight(1 + max(root.Left.GetHeight(), root.Right.GetHeight()))
    
    // balance the tree
    bf := root.BalanceFactor()

    if bf > 1 && val < root.Left.Val {
        return t.RightRotate(root)
    } 
    
    if bf > 1 && val > root.Left.Val {
        root.Left = t.LeftRotate(root.Left)
        return t.RightRotate(root)
    }

    if bf < -1 && val > root.Right.Val {
        return t.LeftRotate(root)
    }
    
    if bf < -1 && val < root.Right.Val {
        root.Right = t.RightRotate(root.Right)
        return t.LeftRotate(root)
    }
    return root
}


type AvlNode struct {
    Val int
    height int
    Left *AvlNode
    Right *AvlNode
}

func NewAvlNode(val int) *AvlNode {
    return &AvlNode{
        Val: val,
        height: 1,
    }
}

func (this *AvlNode) GetHeight() int{
    if this == nil {
        return 0
    }
    return this.height
}

func (this *AvlNode) SetHeight(h int) {
    if this == nil {
        panic("unexpected")
    }
    this.height = h
}

func (this *AvlNode) BalanceFactor() int {
    if this == nil {
        return 0
    }
    return this.Left.GetHeight() - this.Right.GetHeight()
}

func (this *AvlNode) InOrderVal() []int {
    if this == nil {
        return nil
    }

    res := this.Left.InOrderVal()
    res = append(res, this.Val)
    res = append(res, this.Right.InOrderVal()...)
    return res
}

func (this *AvlNode) ToTreeNode() *TreeNode {
    if this == nil {
        return nil
    }

    node := &TreeNode{
        Val: this.Val,
    }
    node.Left = this.Left.ToTreeNode()
    node.Right = this.Right.ToTreeNode()

    return node
}

func sortedListToBST(head *ListNode) *TreeNode {
    avl := &AvlTree{}
    c := head
    for c != nil {
        avl.Insert(c.Val)
        c = c.Next
    }
    // fmt.Println(avl.root.InOrderVal())
    
    return avl.root.ToTreeNode()  
}