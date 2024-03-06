import "container/list"

type LRUCache struct {
    dict map[int]*list.Element
    queue *list.List
    cap int
}

type Pair struct{
    k, v int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        dict: make(map[int]*list.Element),
        queue: list.New(),
        cap: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    el, ok := this.dict[key]

    if !ok {
        return -1
    }

    // move el to the front
    this.queue.MoveToFront(el)
    return el.Value.(Pair).v
}


func (this *LRUCache) Put(key int, value int)  {
    el, ok := this.dict[key]

    if ok {
        this.queue.MoveToFront(el)
        el.Value = Pair{key, value}
        return
    }

    size := len(this.dict)
    if size + 1 > this.cap {
        // remove the back
        back := this.queue.Back()
        this.queue.Remove(back)
        delete(this.dict, back.Value.(Pair).k)
    }
    el = this.queue.PushFront(Pair{key, value})
    this.dict[key] = el
}

func init() { debug.SetGCPercent(-1) }
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */