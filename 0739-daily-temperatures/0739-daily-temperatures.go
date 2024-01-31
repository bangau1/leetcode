func dailyTemperatures(temperatures []int) []int {
    return nge(temperatures)
}

type pair struct {
    index int
    value int
}

func nge(data []int) []int {
    n := len(data)
    stack := make([]pair, 0)
    nge := make([]int, n)
    for i:=0;i<n;i++{
        for len(stack) > 0 && stack[len(stack)-1].value < data[i] {
            pop := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            nge[pop.index] = i - pop.index
        }
        stack = append(stack, pair{i, data[i]})
    }

    return nge
}