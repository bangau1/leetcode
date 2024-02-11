func printGrid(grid [][]int) {
	m := len(grid)
	for i := 0; i < m; i++ {
		fmt.Println(grid[i])
	}

	fmt.Println("=================")
}

func cherryPickup(grid [][]int) int {
	n := len(grid[0])
	return maxCherryPickup([]int{0, 0, n - 1}, grid)
}

type cols struct {
	c1, c2 int
}

func maxCherryPickup(initPos []int, grid [][]int) int {
	row, col1, col2 := initPos[0], initPos[1], initPos[2]
	m, n := len(grid), len(grid[0])
    if m == 1 {
        return grid[row][col1] + grid[row][col2]
    }
	// let dp[r][c1][c2] is the maximum cherry got picked on row=r and c1 and c2 (robot 1 and robot2)
	dp := make([][][]int, m)
    for i:=0;i<m;i++{
        dp[i] = make([][]int, n)
        for ii:=0;ii<n;ii++{
            dp[i][ii] = make([]int, n)
        }
    }
    

    for r:=m-1;r>=row;r--{
        for c1:=0;c1<n;c1++{
            for c2:=0;c2<n;c2++{
                curr := grid[r][c1]
                if c1 != c2 {
                    curr += grid[r][c2]
                }
                subTotal := 0

                if r == m-1 {

                    dp[r][c1][c2] = curr
                    continue
                }
                for nC1 := c1-1; nC1 <= c1+1;nC1++{
                    for nC2:= c2-1;nC2 <= c2+1;nC2++{
                        if nC1 >= 0 && nC1 < n && nC2 >=0 && nC2 < n{
                            subTotal = max(subTotal, dp[r+1][nC1][nC2])
                        }
                    }
                }
                curr += subTotal
                dp[r][c1][c2] = curr
            }
        }
    }

	return dp[row][col1][col2]

}