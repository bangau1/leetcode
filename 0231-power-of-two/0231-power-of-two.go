func isPowerOfTwo(n int) bool {
    c := 0
    for n > 0 {
        n = n & (n-1)
        c++
        if c > 1 {
            return false
        }
    }
    return c == 1
}