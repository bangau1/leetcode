type Stack[T any] struct {
    data []T
}
func NewStack[T any]() Stack[T]{
    return Stack[T]{data:make([]T, 0)}
}

func (this *Stack[T]) Pop() T{
    item := this.data[this.Len()-1]
    this.data = this.data[0:this.Len()-1]
    return item
}

func (this *Stack[T]) Peek() T{
    return this.data[this.Len()-1]
}

func (this *Stack[T]) Push(a T) {
    this.data = append(this.data, a)
}

func (this *Stack[T]) Len() int {
    return len(this.data)
}

func (this *Stack[T]) Empty() bool {
    return this.Len() == 0
}

type MyQueue struct {
    in Stack[int]
    out Stack[int]
}


func Constructor() MyQueue {
    return MyQueue{
        in: NewStack[int](),
        out: NewStack[int](),
    }
}


func (this *MyQueue) Push(x int)  {
    for !this.out.Empty(){
        this.in.Push(this.out.Pop())
    }
    this.in.Push(x)
}


func (this *MyQueue) Pop() int {
    for !this.in.Empty(){
        this.out.Push(this.in.Pop())
    }
    return this.out.Pop()
}


func (this *MyQueue) Peek() int {
    for !this.in.Empty(){
        this.out.Push(this.in.Pop())
    }
    return this.out.Peek()
}


func (this *MyQueue) Empty() bool {
    return this.in.Empty() && this.out.Empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */