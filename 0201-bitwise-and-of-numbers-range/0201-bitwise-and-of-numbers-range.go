func rangeBitwiseAnd(left int, right int) int {
    diff := right - left

    if diff == 0 {
        return right
    }    

    if left == 0 {
        return 0
    }
    res := left
    for i:=left+1;i<=right;i++{
        res = res & i
        if res == 0 {
            return 0
        }
    }
    return res
}

