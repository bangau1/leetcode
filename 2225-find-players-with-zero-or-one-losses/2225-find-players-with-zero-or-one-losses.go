func findWinners(matches [][]int) [][]int {
    loseCount := make(map[int]int)

    for _, match := range matches {
        w, l := match[0], match[1]

        loseCount[l] += 1
        if _, ok:= loseCount[w]; !ok {
            loseCount[w] = 0
        }
    }

    zeroLose := make([]int, 0)
    oneLose := make([]int, 0)

    for user, count := range loseCount {
        if count == 0 {
            zeroLose = append(zeroLose, user)
        }else if count == 1 {
            oneLose = append(oneLose, user)
        }
    }

    sort.Ints(zeroLose)
    sort.Ints(oneLose)
    return [][]int{zeroLose, oneLose}
}