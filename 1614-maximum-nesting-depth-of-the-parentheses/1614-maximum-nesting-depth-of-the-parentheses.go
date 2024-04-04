var open, close byte = '(', ')'

func maxDepth(s string) int {
    var stack int
    maxDepth := 0

    for i:=0;i<len(s);i++{
        if s[i] == open {
            stack++
        }else if s[i] == close {
            if maxDepth < stack {
                maxDepth = stack
            }
            stack--
        }
    }

    return maxDepth
}
