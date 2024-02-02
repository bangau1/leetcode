func sequentialDigits(low int, high int) []int {
    res := make([]int, 0)
    arrs := []int{1,2,3,4,5,6,7,8,9}
    start, end := digitLength(low), min(digitLength(high), 9)
    for digit:=start;digit<=end;digit++{
        // 1 -> 1,2,3,4,5,6,7,8,9
        // 2 -> 12,23,34,45,56,67,78,89
        for i:=0;i<=9-digit;i++{
            num := convert(arrs[i:i+digit])
            if num >= low && num <= high{
                res = append(res, num)
            }
            if num > high {
                break
            }
        }
    }
    return res
}

func digitLength(num int) int {
    len := 0
    for num > 0 {
        len++
        num = num / 10
    }
    return len
}

func convert(arrs []int) int {
    res := 0
    for _, num := range arrs {
        res = res * 10 + num
    }
    return res
}