func removeKdigits(num string, k int) string {
/*
General Idea:
- maintain monotonic stack that's increasing (non-decrease)
- if the stack.peek() element is larger than current digit iteration, pop it while we can (the k > 0)
*/
    
    stack := make([]byte, 0)
    for i:=0;i<len(num);i++{
        for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
            k--
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, num[i])
    }

    // this is to handle the case where k is still positive
    // normally this is because a lot of num digit is equal
    // so remove the last digit while we can
    for k > 0 && len(stack) > 0 {
        k--
        stack = stack[0:len(stack)-1]
    }

    // this is to remove trailing zero
    for len(stack) > 0 && stack[0] == '0' {
        stack = stack[1:]
    }

    // handle empty digit
    if len(stack) == 0 {
        return "0"
    }
    return string(stack)
}