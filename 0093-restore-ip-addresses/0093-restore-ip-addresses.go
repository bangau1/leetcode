var zero = "0"[0]
type result []string

func backtrack(idx int, s string, path []string, answer *result) {
    if idx == len(s) {
        if len(path) == 4 {
            *answer = append(*answer, strings.Join(path, "."))
            // fmt.Println(path)
        }
        return
    }
    if idx > len(s) {
        panic("unexpected")
    }

    if idx < len(s) && len(path) > 4 {
        return
    }

    // to get the segment number, we need to lookup for 1 to 3 digits, each has no 0 leading digits
    prevPathLen := len(path)
    for i:=0;i<3 && idx+i < len(s);i++{
        if i == 0 {
            path = append(path, s[idx:idx+i+1])
            backtrack(idx+i+1, s, path, answer)
            path = path[0:prevPathLen]
        }else{
            if s[idx] != zero {
                num, _ := strconv.Atoi(s[idx:idx+i+1])
                if num <= 255 {
                    path = append(path, s[idx:idx+i+1])
                    backtrack(idx+i+1, s, path, answer)
                    path = path[0:prevPathLen]
                }
            }else{
                break
            }
        }
    }

}


func restoreIpAddresses(s string) []string {
    var answer result
    backtrack(0, s, []string{}, &answer)
    return answer
}