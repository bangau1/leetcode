func numberOfSteps(num int) int {
    c := 0
    for num > 0 {
        if num & 1 == 1 {
            num = (num >> 1 ) << 1
        }else{
            num = num >> 1
        }
        c++
    }
    return c
}