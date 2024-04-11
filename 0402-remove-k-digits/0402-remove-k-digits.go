func removeKdigits(num string, k int) string {
/*
General Idea:
- let n = digit_len(num)
- then we will chose n-k digit from it
- maintain monotonic stack that's increasing (non-decrease)

1432219, k = 3, n = 7
then we need to choose 4 digit
1432 219
1432 2
*/
    stack := make([]byte, 0)
    for i:=0;i<len(num);i++{
        for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
            k--
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, num[i])
    }
    for k > 0 && len(stack) > 0 {
        k--
        stack = stack[0:len(stack)-1]
    }
    for len(stack) > 0 && stack[0] == '0' {
        stack = stack[1:]
    }
    if len(stack) == 0 {
        return "0"
    }
    return string(stack)
}