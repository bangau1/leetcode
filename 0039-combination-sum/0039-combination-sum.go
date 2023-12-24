type result [][]int
// assuming candidates already sorted asc
func generateCombination(candidates []int, target int, start int, answer []int, data *result) {
    if target == 0 {
        *data = append(*data, answer)
        fmt.Println(answer)
        return
    }

    for i:=start;i<len(candidates);i++{
        if candidates[i] <= target {
            newCopy := make([]int, len(answer))
            copy(newCopy, answer)
            newCopy = append(newCopy, candidates[i])
            generateCombination(candidates, target-candidates[i],i, newCopy, data)
        }else{
            break
        }
    }

}

func combinationSum(candidates []int, target int) [][]int {
    data := make(result, 0)
    sort.Ints(candidates)
    generateCombination(candidates, target, 0, []int{}, &data)

    return data
}