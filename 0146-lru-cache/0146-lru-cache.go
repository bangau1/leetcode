type Element[K any, V any] struct{
    Prev *Element[K,V]
    Next *Element[K,V]
    Key K
    Val V
}

func NewElement[K any, V any](k K, v V) Element[K, V] {
    return Element[K,V]{
        Key: k,
        Val: v,
    }
}

func (el *Element[K,V]) addAfter(other *Element[K,V]) {
    other.Next.Prev = el
    el.Prev, el.Next = other, other.Next
    other.Next = el
}

func (el *Element[K,V]) addBefore(other *Element[K,V]) {
    other.Prev.Next = el
    el.Prev, el.Next = other.Prev, other
    other.Prev = el
    
}

type LRUCache struct {
    head, tail *Element[int, int]
    dict map[int]*Element[int,int]
    cap int
}


func Constructor(capacity int) LRUCache {
    h := NewElement[int, int](-1, -1)
    t := NewElement[int, int](-1, -1)
    h.Next, t.Prev = &t, &h
    return LRUCache{
        head: &h,
        tail: &t,
        cap: capacity,
        dict: make(map[int]*Element[int,int]),
    }
}


func (this *LRUCache) Get(key int) int {
    // get from dict[key] -> got the element
    // - if no, return nil
    // - if yes, remove element, and put at the head
    el, ok := this.dict[key]
    if !ok {
        return -1
    }

    // delete element
    this.delete(el)
    // addElement to head
    el.addAfter(this.head)
    this.dict[key] = el

    return el.Val
}


func (this *LRUCache) Put(key int, value int)  {
    // defer func(){
    //     fmt.Println("put", key, value)
    //     this.Print()
    // }()
    // if key exist, update the element
    // if key isn't exists:
    // a. If < capacity, then add to the tail
    // b. if >= capacity, get the tail, update the relevant index
    if el, ok := this.dict[key]; ok {
        el.Val = value
        this.delete(el)
        this.dict[key] = el
        el.addAfter(this.head)
        return
    }
    size := len(this.dict)
    el := NewElement[int, int](key, value)
    if size + 1 > this.cap {
        // get the tail element
        throwEl := this.tail.Prev
        // remove the tail.key from the dict
        this.delete(throwEl)
    }
    // add after head
    this.dict[key] = &el
    el.addAfter(this.head)
}

func (this *LRUCache) Print() {
    curr := this.head.Next
    var res [][2]int
    for curr != this.tail{
        res = append(res, [2]int{curr.Key, curr.Val})
        curr = curr.Next
    }
    fmt.Println(res)
}

func (this *LRUCache) delete(el *Element[int, int]) {
    delete(this.dict, el.Key)
    el.Next.Prev = el.Prev
    el.Prev.Next = el.Next
} 


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */