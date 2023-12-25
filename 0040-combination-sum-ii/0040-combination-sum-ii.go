type result [][]int

func generateCombinationSum(sortedData []int, target int, start int, answer []int, result *result) {
    if target == 0 {
        clone := make([]int, len(answer))
        copy(clone, answer)
        *result = append(*result, clone)
        // fmt.Println(answer)
        return
    }

    n := len(sortedData)
    prevLen := len(answer)
    prevCandidate := -1
    for r := start;r<n;r++{
        if sortedData[r] <= target {
            if sortedData[r] != prevCandidate {
                answer = append(answer, sortedData[r])
                generateCombinationSum(sortedData, target-sortedData[r], r+1, answer, result)
                answer = answer[0:prevLen]
                prevCandidate = sortedData[r]
            }
        }else{ 
            break
        }
    }
}

func combinationSum2(candidates []int, target int) [][]int {
    var result result
    sort.Ints(candidates)
    // fmt.Println(candidates)
    generateCombinationSum(candidates, target, 0, []int{}, &result)
    return result

}
