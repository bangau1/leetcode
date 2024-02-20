func fourSum(nums []int, target int) [][]int {
    n := len(nums)
    if n < 4 {
        return nil
    }

    sort.Ints(nums)
    var res [][]int
    for i:=0;i<n;i++{
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        subRes := threeSum(nums[i+1:], target-nums[i])    
        for _, triSum := range subRes {
            // fmt.Println("3sum", triSum, target-nums[i])
            res = append(res, []int{nums[i], triSum[0], triSum[1], triSum[2]})
        }
    }
    return res
    
}

func threeSum(sorted []int, target int) [][]int {
    n := len(sorted)
    if n < 3 {
        return nil
    }
    var res [][]int
    for i:=0;i<n;i++{
        if i > 0 && sorted[i] == sorted[i-1] {
            continue
        }

        subRes := twoSum(sorted[i+1:], target-sorted[i])    
        for _, twoSumNum := range subRes {
            res = append(res, []int{sorted[i], twoSumNum[0], twoSumNum[1]})
            
        }
    }
    return res
}

func twoSum(sorted []int, target int) [][]int {
    n := len(sorted)

    if n < 2 {
        return nil
    }

    l, r := 0, n-1
    var sum int
    var res [][]int
    for l < r {
        sum = sorted[l] + sorted[r]
        if sum > target {
            r--
            continue
        }else if sum < target {
            l++
            continue
        }else{
            res = append(res, []int{sorted[l], sorted[r]})
            for l + 1 < n && sorted[l+1] == sorted[l] {
                l++
            }
            l++
        }
    }

    return res
}