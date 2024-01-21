func bellmanFord(matrix [][]int) [][]int{
    n := len(matrix)
    
    dist := make([][]int, n)
    for i:=0;i<n;i++{
        dist[i] = make([]int, n)
        for ii:=0;ii<n;ii++{
            dist[i][ii] = matrix[i][ii]
            if i == ii {
                dist[i][ii] = 0
            }
        }
        // fmt.Println(dist[i])
    }
    for k:=0;k<n;k++{
        for u:=0;u<n;u++{
            for v:=0;v<n;v++{
                // dist[u][v] = min(dist[u][v], dist[u][k]+dist[k][u])
                if dist[u][k] != math.MaxInt && dist[k][v] != math.MaxInt{
                    dist[u][v] = min(dist[u][v], dist[u][k]+dist[k][v])
                }
            }
        }
    }
    // fmt.Println("=========")
    
    return dist
    
    
}

func countOfPairs(n int, x int, y int) []int {
    matrix := make([][]int, n)
    
    for i:=0;i<n;i++{
        matrix[i] = make([]int, n)
        for ii:=0;ii<n;ii++{
            matrix[i][ii] = math.MaxInt
            if i == ii {
                matrix[i][ii] = 0
            }
        }
    }
    for i:=0;i<n-1;i++{
        matrix[i][i+1] = 1
        matrix[i+1][i] = 1
    }
    matrix[x-1][y-1] = 1
    matrix[y-1][x-1] = 1
    
    dist := bellmanFord(matrix)
    
    res := make([]int, n)
    for i:=0;i<n;i++{
        // fmt.Println(dist[i])
        for ii:=0;ii<n;ii++{
            if i != ii {
                streetCount := dist[i][ii]
                fmt.Println(streetCount)
                res[streetCount-1]++
            }
        }
    }
    return res
}