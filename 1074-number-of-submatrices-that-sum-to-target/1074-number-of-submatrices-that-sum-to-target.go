func numSubmatrixSumTarget(matrix [][]int, target int) int {
    prefixSum := generatePrefixSum(matrix)
    // for i:=0;i<len(prefixSum);i++{
    //     fmt.Println(prefixSum[i])
    // }

    subRangeSum := func(x1,y1, x2,y2 int) int {
        dp := prefixSum
        return dp[x2+1][y2+1] - dp[x1][y2+1] - dp[x2+1][y1] + dp[x1][y1]
    }
    m, n := len(matrix), len(matrix[0])
    total := 0
    for r1:=0;r1<m;r1++{
        for r2:=r1;r2<m;r2++{
            // we want to generate the 1D prefix sum
            oneSums := make([]int, n+1)
            for c:=0;c<n;c++{
                oneSums[c+1] = subRangeSum(r1,0,r2,c)
            }
            total += subarraySum(oneSums, target)
        }
    }
    
    return total
}

func generatePrefixSum(matrix [][]int) [][]int {
    m, n := len(matrix), len(matrix[0])

    dp := make([][]int, m+1)
    for i:=0;i<m+1;i++{
        dp[i] = make([]int, n+1)
    }

    for i:=1;i<=m;i++{
        for ii:=1;ii<=n;ii++{
            dp[i][ii] = matrix[i-1][ii-1] + dp[i-1][ii] + dp[i][ii-1] - dp[i-1][ii-1]
        }
    }
    return dp
}

func subarraySum(prefixSum []int, target int) int {
    n := len(prefixSum)

    counter := make(map[int]int)
    counter[0] = 1
    count := 0
    for i:=1;i<n;i++{
        sum := prefixSum[i]
        count += counter[sum-target]
        counter[sum]++
    }
    return count
}