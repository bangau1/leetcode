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
    var minSum = math.MaxInt
    calculateMinSum(0, workers, bikes, 0, 0, &minSum)
    return minSum
}

