func minimumCost(source string, target string, original []byte, changed []byte, costs []int) int64 {
    // Now, based on the rules (origin -> changed, with the cost), we need to precompute the minimum cost to change character A to B
    // we can use djikstra with all source
    //
    // After that we can calculate the minimum edit distance from source to target
    costMatrix := shortestDistance(original, changed, costs)

    // for i:=0;i<26;i++{
    //     fmt.Println(costMatrix[i])
    // }
    n := len(source)
    res := int64(0)
    for i:=0;i<n;i++{
        from := int(source[i]-'a')
        to := int(target[i]-'a')
        if costMatrix[from][to] != -1 {
            
            res = res + int64(costMatrix[from][to])
        }else{
            return -1
        }
    }
    return res
}

func shortestDistance(original []byte, changed []byte, costs []int) [26][26]int{
    matrix := [26][26]int{}
    for i:=0;i<26;i++{
        for ii:=0;ii<26;ii++{
            matrix[i][ii] = -1
        }
        matrix[i][i] = 0
    }

    n := len(original)
    for i:=0;i<n;i++{
        src, dst, cost := int(original[i]-'a'), int(changed[i]-'a'), costs[i]
        if matrix[src][dst] == -1 {
            matrix[src][dst] = cost
        }else{
            matrix[src][dst] = min(matrix[src][dst], cost)
        }
        
    }

    return floydWarshall(matrix)
}

func floydWarshall(matrix [26][26]int) [26][26]int {
    for k:=0;k<26;k++{
        for i:=0;i<26;i++{
            for j:=0;j<26;j++{
                if matrix[i][k] == -1 || matrix[k][j] == -1 {
                    continue
                }
                if matrix[i][j] == -1 {
                    matrix[i][j] = matrix[i][k] + matrix[k][j]
                }else{
                    matrix[i][j] = min(matrix[i][j],matrix[i][k] + matrix[k][j])
                }
            }
        }
    }
    return matrix
}