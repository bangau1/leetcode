func maximumOddBinaryNumber(s string) string {
    res := []byte(s)
    n := len(res)
    zero := byte('0')
    one := byte('1')
    l := n-2
    if res[n-1] == zero {
        for l >= 0 && res[l] == zero{
            l--
        }
        res[l], res[n-1] = res[n-1], res[l]
    }
    c := 0
    for i:=0;i<=l;i++{
        if res[i] == one {
            c++
        }
        res[i] = zero
    }
    for i:=0;i<c;i++{
        res[i] = one
    }
    return string(res)

}