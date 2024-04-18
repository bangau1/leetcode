var dirs = [][2]int{
    {-1,0},
    {1,0},
    {0, -1},
    {0, 1},
}

func islandPerimeter(grid [][]int) int {
    sum := 0
    nr, nc := 0, 0
    neighbor := 0
    for r:=0;r<len(grid);r++{
        for c:=0;c<len(grid[r]);c++{
            if grid[r][c] == 1 {
                neighbor  = 0
                for _, dir := range dirs {
                    nr, nc = r+dir[0], c+dir[1]
                    if nr >= 0 && nc >= 0 && nr < len(grid) && nc < len(grid[r]) && grid[nr][nc] == 1 {
                        neighbor += 1
                    }
                }
                sum += 4-neighbor
            }
        }
    }
    return sum
}