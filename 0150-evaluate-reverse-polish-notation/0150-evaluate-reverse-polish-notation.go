func evalRPN(tokens []string) int {
    if len(tokens) == 1{
        return mustNum(tokens[0])
    }
    l := len(tokens)
    stack := make([]string, 0)
    i := l-1
    for i >= 0 {
        
        for shouldEvaluate(stack) {
            // fmt.Println(stack)
            sl := len(stack)

            a, b := mustNum(stack[sl-1]), mustNum(stack[sl-2])
            num := calculate(a, b, stack[sl-3])
            
            stack = stack[:sl-3]
            stack = append(stack, fmt.Sprintf("%d", num))
            
        }
        stack = append(stack, tokens[i])
        // fmt.Println(stack)
        i--
    }
    for shouldEvaluate(stack) {
        // fmt.Println(stack)
        sl := len(stack)

        a, b := mustNum(stack[sl-1]), mustNum(stack[sl-2])
        num := calculate(a, b, stack[sl-3])
        
        stack = stack[:sl-3]
        stack = append(stack, fmt.Sprintf("%d", num))
        
    }
    return mustNum(stack[0])
}
func shouldEvaluate(data []string) bool {
    l := len(data)
    if l < 3 {
        return false
    }
    a, b, operand := data[l-1], data[l-2], data[l-3]
    aNum := isNumber(a)
    bNum := isNumber(b)
    operandNum := isNumber(operand)
    return aNum && bNum && !operandNum
}
func mustNum(a string) int {
    num, err := strconv.Atoi(a)
    if err != nil {
        fmt.Println(a, "is not num")
        panic("error")
    }
    return num
}

func calculate(a, b int, token string) int {
    switch token {
        case "+":
        a = a + b
        case "-":
        a = a - b
        case "*":
        a = a * b
        case "/":
        a = a /b
        default:
        // no-op
    }
    return a
}

func isNumber(r string) bool {
    return !isOperand(r)
}

func isOperand(r string) bool {
    return len(r) == 1 && (r == "-" || r == "*" || r == "/" || r == "+") 
}