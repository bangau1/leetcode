func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    if a == 0 {
        return b
    }
    if a < b {
        a, b = b, a
    }
    return gcd(b, a % b)
}

func originalSolution(nums []int, k int) int {
    if k == 0 {
        return 0
    }

    n:=len(nums)

    sol := make([][]int, 0)
    left, right:=0,0
    dot := 1
    prevDot := 1
    
    for left < n && right < n {
        prevDot = dot
        dot *= nums[right]
        // fmt.Println("dot", dot, "from", nums[left:right+1])

        if dot < k {
            right++
        }else{
            if prevDot < k {
                sol = append(sol, []int{left, right-1})
            }
            for dot >= k && left <= right{
                dot /= nums[left]
                left++
            }
            right++
        }
    }
    if prevDot < k {
        sol = append(sol, []int{left, right-1})
    }
    // fmt.Println(sol)
    var total int
    for i:=0;i<len(sol);i++{
        left, right := sol[i][0], sol[i][1]
        arrLen := (right-left+1)
        combi := 0
        if arrLen % 2 == 0{
            combi = arrLen/2 * (arrLen+1)
        }else{
            combi = (arrLen+1)/2 * arrLen
        }

        if i > 0 {
            // check overlap range
            _, prevRight := sol[i-1][0], sol[i-1][1]
            // overlap
            if prevRight >= left {
                overlap := prevRight-left+1
                combi -= (overlap*(overlap+1))/2 
            }
        }
        // fmt.Println(sol[i], combi)
        total+=combi
    }
    return total
}

func numSubarrayProductLessThanK(nums []int, k int) int {
    return simpleSolution(nums, k)
}

func simpleSolution(nums []int, k int) int {
    if k <= 1 {
        return 0
    }

    n:=len(nums)
    l := 0
    dot := 1
    total := 0
    for r:=0;r<n;r++{
        dot *= nums[r]

        for dot >= k {
            dot /= nums[l]
            l++
        }
        total += r-l +1
    }

    return total
}


// a, b, c, d
// if a..c and b...d is the max subarray we found, then 
// the number of subarray are:
// a,b,c -> there are n * n+1/2 = 6
// b,c,d -> there are 6
// overlap b,c ->3 combination 
// total 6+6-3=9


// 10,5 = 3 combi
// 5,2,6 = 6 combi
// 5 = 1 combi
// 9-1=8

// 1 1 1
//  2. 2
//.   3