func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    // ["practice", "makes", "perfect", "coding", "makes"], word1 = "makes", word2 = "coding"
    // makes = [1,4]
    // coding = [3,6]
    // 1,3 = 2
    // 3,4 = 1
    // 

    // naive solution
    // a := make([]int, 0)
    // b := make([]int, 0)
    // for i, word := range wordsDict {
    //     if word == word1 {
    //         a = append(a, i)
    //     }else if word == word2{
    //         b = append(b, i)
    //     }
    // }
    // iA, iB := 0, 0
    // minRes := abs(a[iA]-b[iB])
    
    // var diff int
    // for iA < len(a) && iB < len(b) {
    //     diff = abs(a[iA]-b[iB])

    //     for iA < len(a) && iB < len(b) && a[iA] < b[iB] {
    //         diff = min(diff, b[iB] - a[iA])
    //         iA++
    //     } 

    //     for iA < len(a) && iB < len(b) && b[iB] < a[iA] {
    //         diff = min(diff, a[iA] - b[iB])
    //         iB++
    //     }
    //     minRes = min(minRes, diff)
    // }

    // optimize one pass solution
    var minRes = math.MaxInt
    i, j := -1, -1

    for idx, word := range wordsDict {
        if word == word1{
            i = idx
        }else if word == word2 {
            j = idx
        }

        if i != -1 && j != -1 {
            minRes = min(minRes, abs(i-j))
        }
    }

    return minRes
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}