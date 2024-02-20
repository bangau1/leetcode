func divide(dividend int, divisor int) int {
    res := doDivide(dividend, divisor)
    if res < -2147483648 {
        return -2147483648
    }
    if res > 2147483647 {
        return 2147483647
    }
    return res
}
func doDivide(dividend int, divisor int) int {
    if divisor == 1 {
        return dividend
    }

    // 10 = 1010
    // 2  = 0010
    // 5  = 0101

    // 10 = 1010
    // 3  = 0011
    // 5  = 0011

    if dividend < 0 && divisor < 0 {
        return divide(-dividend, -divisor)
    }
    if divisor < 0 {
        return -divide(dividend, -divisor)
    }
    if dividend < 0 {
        return -divide(-dividend, divisor)
    }
   
    if dividend == divisor {
        return 1
    } 
    
    if dividend < divisor {
        return 0
    }
    // a = b*(2^N) + b *(2^M)
    // a/b = 2^N + 2 ^M
    a, b := dividend, divisor
    i := 0
    for (b << i) < a {
        i++
    }
    if b << i == a {
        return 1 << i
    }
    i = i -1
    res := 1 << i
    left := a - (b << i)
    for i >= 0 && left > 0 {
        for i >= 1 && b << (i-1) > left {
            i--
        }
        if left >= (b <<i){
            res += 1 << i
            left -= b << i
        }else {
            i--
        }
    }

    return res
}