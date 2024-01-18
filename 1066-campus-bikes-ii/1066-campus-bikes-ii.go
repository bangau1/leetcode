func calculateMinSum(worker int, workers [][]int, bikes [][]int, bikeFlag int, currentSum int, minSum *int) {
    n := len(workers)
    m := len(bikes)

    if worker == n {
        if *minSum > currentSum {
            *minSum = currentSum
        }
        return
    }

    if currentSum > *minSum {
        return
    }

    for i:=0;i<m;i++{
        if bikeFlag & (1<<i) == 0{
            prevFlag := bikeFlag
            bikeFlag = bikeFlag | (1 << i)
            calculateMinSum(worker+1, workers, bikes, bikeFlag, currentSum + dist(workers[worker], bikes[i]), minSum)
            bikeFlag = prevFlag
        }
    }
}

// kernighan algo
func countSetBit(n int) int {
    count := 0
    for n != 0 {
        n &= (n-1)
        count++
    }
    return count
}
func dpSolution(workers [][]int, bikes [][]int) int {

    dp := make([]int, 1 << len(bikes))
    for i:=0;i<len(dp);i++{
        dp[i] = math.MaxInt
    }
    dp[0] = 0 // no bike being assigned, so the sum is 0
    minSolution := math.MaxInt

    for mask := 0;mask < 1 << len(bikes);mask++{
        nextWorkerIdx := countSetBit(mask)

        if nextWorkerIdx >= len(workers) {
            minSolution = min(minSolution, dp[mask])
            continue
        }

        for i:=0;i<len(bikes);i++{
            
            if mask & (1 << i) == 0 {
                newMask := mask | 1 << i
                dist := dist(workers[nextWorkerIdx], bikes[i])
                dp[newMask] = min(dp[newMask], dp[mask] + dist)
            }
        }


    }
    return minSolution

}


func dist(workerPos []int, bikePos []int) int {
    return abs(workerPos[0]-bikePos[0]) + abs(workerPos[1] - bikePos[1])
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func assignBikes(workers [][]int, bikes [][]int) int {
    // var minSum = math.MaxInt
    // calculateMinSum(0, workers, bikes, 0, 0, &minSum)
    // return minSum

    return dpSolution(workers, bikes)
}

