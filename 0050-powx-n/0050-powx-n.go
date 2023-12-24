func abs (a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func myPow(x float64, n int) float64 {
    if x == float64(1) || x == float64(0){
        return x
    }
    if n == 0 {
        return float64(1)
    }


    res := x
    loop := abs(n)
    leftLoop := 0
    if loop > 1 {
        for i:=2;i<=loop;i+=i{
            res *= res
            if i+i > loop {
                leftLoop = loop - i
                break
            }
        }
        for i:=1;i<=leftLoop;i++{
            res *= x
        }
    }
    if n < 0 {
        res = float64(1)/res
    }
    
    return res
}