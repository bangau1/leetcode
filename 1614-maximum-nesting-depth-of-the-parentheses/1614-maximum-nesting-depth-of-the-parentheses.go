var open, close byte = '(', ')'

func maxDepth(s string) int {
    stack := make([]byte, 0)
    maxDepth := 0

    for i:=0;i<len(s);i++{
        if s[i] == open {
            stack = append(stack, s[i])
        }else if s[i] == close {
            if maxDepth < len(stack) {
                maxDepth = len(stack)
            }
            stack = stack[:len(stack)-1]
        }
    }

    return maxDepth
}
