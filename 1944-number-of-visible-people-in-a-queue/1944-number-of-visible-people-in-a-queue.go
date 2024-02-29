func canSeePersonsCount(h []int) []int {
    // the general idea is to maintain monotonous stack in decreasing manner from right iteratively
    // - let say we have maintain monotonous stack until index i
    // - then people at i-1, they can only see in the stack until the people in stack is greater than it (no equal since unique constraint)
    stack := make([]int, 0)
    n := len(h)
    res := make([]int, n)
    
    for i:=n-1;i>=0;i--{
        
        for len(stack) > 0 && h[i] > h[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
            res[i]++
        }
        if len(stack) > 0 {
            res[i]+=1
        }
        stack = append(stack, i)
    }
    return res
}