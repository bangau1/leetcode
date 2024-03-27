func getGCD(a, b int) int {
    if a < b {
        a, b = b, a
    }

    if b == 0 {
        return a
    }

    return getGCD(b, a % b)
}


func numSubarrayProductLessThanK(nums []int, k int) int {
    // since len(nums) is quite big, the product of all elements in the subarray can be really large, doesn't fit in the int64
    // so one thing that we can optimize is to divide the product of all element by the gcd(k, product(arr))
    // assume there is subarray of [a, b, c], to check whether a*b*c < k, we can divide both by the gcd in iterative manner:
    // - a/gcd(a,k) * b/gcd(b, k') * c /gcd(c, k'') < k''

    // now how to compute how many subarray on that? we can use sliding window that will expand as long as the condition true.
    // if not satisfied, we decrease the window 
    // everytime we increase the window length, then that increase the total subset += len(window)

    r := 0

    var gcd int
    var prod = 1
    var sum = 0

    window := make([][2]int, 0)
    for r < len(nums) {
        gcd = getGCD(nums[r], k)
        // gcd = 1
        prod = prod * nums[r]/gcd
        k = k /gcd

        window = append(window, [2]int{nums[r], gcd})
        
        if prod < k {
            sum += len(window)
        }else{
            for len(window) > 0 && prod >= k {

                prod = prod * window[0][1]/ window[0][0]
                k = k * window[0][1]
                window = window[1:]
            }

            sum += len(window)
        }
        r++
    }

    return sum
}