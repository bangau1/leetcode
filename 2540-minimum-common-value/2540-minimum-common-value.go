func getCommon(nums1 []int, nums2 []int) int {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    ln, rn := len(nums1), len(nums2)
    i := sort.SearchInts(nums1, nums2[0])
    var idx int
    for ;i<ln;i++{
        idx = sort.SearchInts(nums2, nums1[i])
        if idx < rn && nums2[idx] == nums1[i] {
            return nums1[i]
        }
    }

    return -1
}