
func twoSum(nums []int, target int) []int {
    sumMap := make(map[int]int)

    for i, num := range nums {
        sumMap[num] = i // if there's same num occuring then the latest i will be recorded
    }

    for i, num := range nums {
        diff := target - num
        if _, ok:=sumMap[diff];ok && sumMap[diff] != i {
            return []int{i, sumMap[diff]}
        }
    }

    return []int{-1, -1}
}