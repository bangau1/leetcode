func getCommon(nums1 []int, nums2 []int) int {
    ln, rn := len(nums1), len(nums2)
    l, r := 0, 0

    for l < ln && r < rn {
        if nums1[l] < nums2[r] {
            l++
        }else if nums1[l] > nums2[r] {
            r++
        }else{
            return nums1[l]
        }
    }

    return -1
}