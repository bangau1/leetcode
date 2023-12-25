type result [][]int

func backtrack(start int, k int, target int, path []int, answer *result) {
    if target == 0 {
        if len(path) == k {
            combi := make([]int, len(path))
            copy(combi, path)
            *answer = append(*answer, combi)    
        }
        return
    }

    lenPath := len(path)
    for next:=start;next<=9;next++{
        if next <= target {
            path = append(path, next)
            backtrack(next+1, k, target-next, path, answer)
            path = path[0:lenPath]
        }else{
            break
        }
    }
}

func combinationSum3(k int, n int) [][]int {
    var answer result
    backtrack(1, k, n, []int{}, &answer)
    return answer
}