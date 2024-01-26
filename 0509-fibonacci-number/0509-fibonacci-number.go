
func fib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    n2 := 0
    n1 := 1
    res := 0
    for i:=2;i<=n;i++{
        res = n2 + n1
        n2 = n1
        n1 = res
    }   

    return res
}