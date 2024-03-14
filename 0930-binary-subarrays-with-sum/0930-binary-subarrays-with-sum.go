func numSubarraysWithSum(nums []int, goal int) int {
    // binary array = either 1 or 0
    // calculate total non-empty subarrays with sum = goal

    // general idea:
    // - it can be solved by using a prefix sum + hash table lookup
    n := len(nums)
    hashMap := make(map[int]int)
    hashMap[0] = 1 // the prefixSum at the empty set
    var search int
    var total int
    var prefixSum int
    for i:=1; i<=n; i++ {
        prefixSum += nums[i-1]
        search  = prefixSum - goal
        if hashMap[search] > 0 {
            total+=hashMap[search]
        }
        hashMap[prefixSum] += 1
    }
    return total
}