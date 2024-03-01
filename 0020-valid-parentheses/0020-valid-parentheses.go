
func isValid(s string) bool {
    stack := make([]byte, 0)
    for i:=0;i<len(s);i++{
        if isOpen(s[i]){
            stack = append(stack, s[i])
            continue
        }
        // s[i] is a close parentheses
        if len(stack) > 0 && stack[len(stack)-1] == getOpen(s[i]){
            stack = stack[:len(stack)-1]
            continue
        }
        return false

    }
    return len(stack) == 0
}

func isOpen(a byte) bool {
    return a == '(' || a == '{' || a == '['
}

func getOpen(close byte) byte {
    if close == ')' {
        return '('
    }
    if close == '}' {
        return '{'
    }
    if close == ']' {
        return '['
    }
    panic("unexpected")
}