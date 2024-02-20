type three [3]int

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    dict := make(map[three]bool)
    n := len(nums)
    var data []int
    for i:=0;i<n;i++{
        twoRes := twoSum(nums[i+1:], -nums[i])
        for _, twoSums := range twoRes {
            data = []int{nums[i], twoSums[0], twoSums[1]}
            sort.Ints(data)
            dict[three{data[0], data[1], data[2]}] = true
        }
    }

    for key, _ := range dict {
        res = append(res, []int{key[0], key[1], key[2]})
    }

    return res
}

func twoSum(sorted []int, target int) [][]int {
    n := len(sorted)
    if n < 2 {
        return nil
    }
    var res [][]int
    for i:=0;i<n;i++{
        search := target-sorted[i]

        idx := sort.SearchInts(sorted[i+1:], search)
        if idx+i+1 < n && sorted[i+1+idx] == search {
            res = append(res, []int{sorted[i], search})
        }
    }

    return res
}