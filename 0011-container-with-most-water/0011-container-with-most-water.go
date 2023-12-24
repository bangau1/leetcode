func maxArea(height []int) int {
    n := len(height)
    // use two pointer approach
    l, r := 0, n-1

    area := math.MinInt
    for l < r {
        curArea := (r-l) * min(height[l], height[r])
        area = max(area, curArea)
        if height[l] > height[r] {
            r--
        }else if height[r] > height[l] {
            l++
        }else{
            // so pick one
            l++
        }
    }
    return area
}