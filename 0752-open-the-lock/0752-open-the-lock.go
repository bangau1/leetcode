
type Wheels struct {
	data [4]byte
}

func (w *Wheels) String() string {
	if w == nil {
		return ""
	}
	res := make([]byte, 4)
	res[0] = w.data[0]
	res[1] = w.data[1]
	res[2] = w.data[2]
	res[3] = w.data[3]
	return string(res)
}

type Vertex struct {
	w Wheels
	c int
}

func (v *Vertex) getEstCost(target Wheels) int {
	cost := 0
	var minCost = 0
	// 0 -> 7 can be: 0->1->2->...->7 = 7 steps, or 0->9->8->7 = 3 steps
	// 9->7
	for i := 0; i < 4; i++ {
		left, right := v.w.data[i], target.data[i]
		if left == right {
			continue
		}

		if left > right {
			left, right = right, left
		}

		minCost = min(
			int(right-left),
			10-int(right+left),
		)

		cost += minCost
	}

	return cost
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func NewWheels(str string) Wheels {
	return Wheels{
		data: [4]byte{str[0], str[1], str[2], str[3]},
	}
}

func NewWheelsFromWheels(other Wheels) Wheels {
	return Wheels{
		data: [4]byte{other.data[0], other.data[1], other.data[2], other.data[3]},
	}
}

func (w *Wheels) getNextSteps() []Wheels {

	res := make([]Wheels, 0)
	for i := 0; i < 4; i++ {
		currDigit := int(w.data[i] - '0')
		currDigit = (currDigit + 1) % 10
		newWheel := NewWheelsFromWheels(*w)
		newWheel.data[i] = byte(currDigit) + '0'
		res = append(res, newWheel)

		currDigit = int(w.data[i] - '0')
		currDigit = currDigit - 1
		if currDigit < 0 {
			currDigit += 10
		}

		newWheel = NewWheelsFromWheels(*w)
		newWheel.data[i] = byte(currDigit) + '0'
		res = append(res, newWheel)
	}
	return res
}

func openLock(deadends []string, target string) int {
	ignore := make(map[Wheels]bool)

	for _, d := range deadends {
		w := NewWheels(d)
		ignore[w] = true
	}

	visited := make(map[Wheels]bool)
	s := NewWheels("0000")
	t := NewWheels(target)
	q := NewMinHeap[Vertex](func(a, b Vertex) bool {
        if a.c == b.c {
            return a.getEstCost(t) < b.getEstCost(t)
        }
		return a.c < b.c
	},
		Vertex{s, 0},
	)

	if ignore[t] || ignore[s] {
		return -1
	}

	var node Vertex
	for q.Len() > 0 {
		node = heap.Pop(q).(Vertex)

		if visited[node.w] {
			continue
		}

		visited[node.w] = true
		if node.w == t {
			return node.c
		}

		for _, next := range node.w.getNextSteps() {
			if !visited[next] && !ignore[next] {
				heap.Push(q, Vertex{next, node.c + 1})
			}
		}
	}
	return -1
}


type MinHeap[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewMinHeap[T any](less func(a, b T) bool, data ...T) *MinHeap[T] {
	res := MinHeap[T]{
		less: less,
	}

	if len(data) > 0 {
		res.data = make([]T, len(data))
		copy(res.data, data)
		heap.Init(&res)
	}

	return &res
}

func (m MinHeap[T]) Less(a, b int) bool {
	return m.less(m.data[a], m.data[b])
}

func (m MinHeap[T]) Swap(a, b int) {
	m.data[a], m.data[b] = m.data[b], m.data[a]
}

func (m MinHeap[T]) Len() int {
	return len(m.data)
}

func (m *MinHeap[T]) Pop() any {
	l := m.Len()
	item := m.data[l-1]
	m.data = m.data[0 : l-1]
	return item
}

func (m *MinHeap[T]) Push(a any) {
	m.data = append(m.data, a.(T))
}
