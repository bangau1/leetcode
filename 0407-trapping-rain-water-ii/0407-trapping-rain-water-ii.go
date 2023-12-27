type data struct {
    x, y int
    h int
}

type minHeap []data

func (m minHeap) Less(a, b int) bool {
    // we want to return the minimum one (law of minimum)
    return m[a].h < m[b].h
}

func (m minHeap) Len() int {
    return len(m)
}

func (m minHeap) Swap(a, b int) {
    m[a], m[b] = m[b], m[a]
}

func (m *minHeap) Pop() any {
    l := m.Len()
    item := (*m)[l-1]
    *m = (*m)[0:l-1]
    return item
}
func (m *minHeap) Push(a any) {
    *m = append(*m, a.(data))
}

func trapRainWater(nums [][]int) int {
    m,n := len(nums), len(nums[0])
    if m < 3 || n < 3 {
        return 0
    }

    pq := minHeap{}
    visited := make([][]bool, m)
    for i:=0;i<m;i++{
        visited[i] = make([]bool, n)
    }

    // add initial boundary (the outer rectangle)   
    for i:=0;i<m;i++{
        // outer left and right 
        pq = append(pq, data{i, 0, nums[i][0]})
        pq = append(pq, data{i, n-1, nums[i][n-1]})
        visited[i][0] = true
        visited[i][n-1] = true
    }

    for i:=1;i<n-1;i++{
        // top and bottom outer rectangle
        pq = append(pq, data{0, i, nums[0][i]})
        pq = append(pq, data{m-1, i, nums[m-1][i]})
        visited[0][i] = true
        visited[m-1][i] = true
    }
    dirs := [][]int{
        {-1,0},
        {1, 0},
        {0, -1},
        {0, 1},
    }
    heap.Init(&pq)
    total := 0
    for pq.Len() > 0 {
        node := heap.Pop(&pq).(data)

        for _, dir := range dirs {
            x, y := node.x + dir[0], node.y + dir[1]
            
            if x < 0 || y < 0 || x >= m || y >= n || visited[x][y] {
                continue
            }
            //new boundary is the maximum between current minData or its neighbor
            h := max(node.h, nums[x][y])

            if nums[x][y] < node.h {
                total+= node.h - nums[x][y]
            }
            visited[x][y] = true
            heap.Push(&pq, data{x, y, h})
        }
    }

    return total
}