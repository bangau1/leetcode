var track = make([]int, 26)

func longestIdealString(s string, k int) int {
/*
Bruteforce approach:
- backtrack, exhaust on all possible combination, to choose s[i] or not.
- this is not feasible as it has 2^n time complexity, where n = len(s)

observations:
- let dp[i] = longest ideal subsequence that ends with s[i]
dp[i]
- at index i, 
*/

/*
- k=2 acfgbd -> ac[f][g]bd
- k=2 dcfg
- k=2 acfghi
*/

    clearArr[int](track, 0)
    res := 1
    for i:=0;i<len(s);i++{
        c := int(s[i]-'a')
        currMax := track[c]
        left, right := max(0, c-k), min(25, c+k)
        for l := left; l<=right;l++{
            if currMax < track[l]{
                currMax = track[l]
            }
            if currMax == i {
                break
            }
        }
        track[c] = currMax + 1
        if res < track[c] {
            res = track[c]
        }
    }
    return res
}

func clearArr[T any](arr []T, val T) {
    for i:=0;i<len(arr);i++{
        arr[i] = val
    }
}