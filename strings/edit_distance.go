package strings

func EditDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)

	if m == 0 {
		return n
	} else if n == 0 {
		return m
	} else if m == n && s1 == s2 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for ii := 1; ii <= n; ii++ {
			if s1[i-1] == s2[ii-1] {
				dp[i][ii] = dp[i-1][ii-1]
			} else {
				dp[i][ii] = min(
					dp[i-1][ii],   // delete operation
					dp[i][ii-1],   // insert operation
					dp[i-1][ii-1], // 1 delete operation
				) + 1
			}
		}
	}

	return dp[m][n]

}
