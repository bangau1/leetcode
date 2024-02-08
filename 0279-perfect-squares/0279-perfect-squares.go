var psq = make([]int, 0)

var dp []int
func init() {
    for i:=1; i <= 100;i++{
        sq := i*i
        psq = append(psq, sq)
    }
    n := 10000
    dp = make([]int, n+1)
    for i:=0;i<=n;i++{
        dp[i] = -1
    }
    dp[0] = 0
    
    for i:=1;i<=n;i++{
        if dp[i] == 1 {
            continue
        }

        // looking the greatest prevSquare < i
        upper := sort.Search(len(psq), func(idx int) bool {
            return psq[idx] > i
        })
        res := math.MaxInt
        for ii:= upper-1;ii>=0 ;ii--{
            res = min(res, 1 + dp[i-psq[ii]])
        }
        dp[i] = res
    }

}

func numSquares(n int) int {
    return dp[n]
}