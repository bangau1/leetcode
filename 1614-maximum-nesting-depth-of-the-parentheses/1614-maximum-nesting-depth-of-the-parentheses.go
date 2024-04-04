var open, close byte = '(', ')'
var stack, maxD int
var n int

func maxDepth(s string) int {
    stack, maxD = 0, 0
    n = len(s)
    for i:=0;i<n;i++{
        if s[i] == open {
            stack++
        }else if s[i] == close {
            if maxD < stack {
                maxD = stack
            }
            stack--
        }
    }

    return maxD
}
