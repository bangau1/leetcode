func minOperations(s string) int {
    // do 2 passes, and see which one is minimum
    even := "0"[0]
    odd := "1"[0]
    ops := 0
    for i:=0;i<len(s);i++{
        if i % 2 == 0 && s[i] != even{
            ops++
        }else if i%2 == 1 && s[i] != odd{
            ops++
        }
    }
    minOps := ops
    even = "1"[0]
    odd = "0"[0]
    ops = 0
    for i:=0;i<len(s);i++{
        if i % 2 == 0 && s[i] != even{
            ops++
        }else if i%2 == 1 && s[i] != odd{
            ops++
        }
    }
    return min(minOps, ops)
}