func addRungs(rungs []int, dist int) int {
    n := len(rungs)
    
    total := 0
    for i:=0;i<n;i++{
        diff := 0
        if i > 0 {
            diff = rungs[i] - rungs[i-1]
        }else{
            diff = rungs[i]
        }
        if diff > dist {
            if diff % dist == 0 {
                total += diff/dist - 1
            }else{
                total += diff/dist
            }
        }
    }
    return total
}