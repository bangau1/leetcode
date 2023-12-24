func commonPrefix(a, b string) string {
    // set a always smaller (or eq) in length than b
    if len(a) > len(b) {
        a, b = b, a
    }

    for i:=0;i<len(a);i++{
        if a[i] != b[i]{
            return a[0:i]
        }

    }
    return a
}

func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }

    common := strs[0]
    for i:=1;i<len(strs);i++{
        common = commonPrefix(common, strs[i])
    }

    return common
}