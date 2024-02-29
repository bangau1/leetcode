func canSeePersonsCount(h []int) []int {
    // the general idea is to maintain monotonous stack in decreasing manner from right iteratively
    // - let say we have maintain monotonous stack until index i
    // - then people at i-1, they can only see in the stack until the people in stack is greater than it (no equal since unique constraint)
    stack := make([]int, 0)
    n := len(h)
    res := make([]int, n)
    
    last := n-1
    for i:=n-2;i>=0;i--{
        for ii:=last;ii>=i+1;ii--{
            for len(stack) > 0 && h[ii] > h[stack[len(stack)-1]] {
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, ii)
        }

        last = i
        if h[i] > h[stack[0]] {
            res[i] = len(stack)
            continue
        }
        // look at the smallest idx, such that h[stack[idx]] > h[i]
        // then the number of people is idx+1
        // but our stack is reverse actually
        //  [8,5,1] 3 => 2 people
        //  search left idx that is < 
        idx := sort.Search(len(stack), func(idx int) bool {
            return h[stack[idx]] < h[i]
        })

        if idx-1 >= 0 {
            res[i] = len(stack)-(idx-1)
        }else{
            res[i] = len(stack)
        }
    }
    return res
}