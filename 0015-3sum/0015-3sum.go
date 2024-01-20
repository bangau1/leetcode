type pair struct{
    a, b, c int
}
func twoSum(counting map[int]int, nums []int, start, end, target int) [][]int {
    res := make([][]int, 0)
    for num, count := range counting {
        if  target - num == num  && count >= 2 {
            res = append(res, []int{num, num})
        }else if target-num != num && count >= 1 && counting[target-num] >= 1{
            res = append(res, []int{num, target-num})
        }
    }
    // for i:=start;i<=end;i++{
    //     search := target - nums[i]

    //     counting[nums[i]]--
    //     if counting[search] > 0 {
    //         res = append(res, []int{nums[i], search})
    //     }

    //     counting[nums[i]]++
    // }
    return res
}

func threeSum(nums []int) [][]int {
    counting := make(map[int]int)
    for _, num := range nums {
        counting[num]++
    }

    flag := make(map[pair]bool)
    res := make([][]int, 0)
    for i:=0;i<len(nums);i++{
        counting[nums[i]]--
        if counting[nums[i]] == 0 {
            delete(counting, nums[i])
        }
        
        data := twoSum(counting, nums, i+1, len(nums)-1, -nums[i])
        
        for _, datum := range data {
            arr := []int{nums[i], datum[0], datum[1]}
            sort.Ints(arr)
            p := pair{arr[0], arr[1], arr[2]}
            if !flag[p] {
                res = append(res, arr)
                flag[p] = true
            }
        }
    }
    return res
}