func strStr(haystack string, needle string) int {
    m, n := len(haystack), len(needle)
    if n > m {
        return -1
    }
    // O(n*m) = naive solution
    // for i:=0;i<m-n+1;i++{
    //     if needle == haystack[i:i+n]{
    //         return i
    //     }
    // }
    // return -1

    // kmp solution o(n+m)
    return kmpSolution(haystack, needle)
}

func kmpPrefixTable(haystack, needle string) []int {
    n:= len(needle)
    lsp := make([]int, n)
    len, i := 0, 1

    for i < n {
        if needle[i] == needle[len] {
            lsp[i] = len + 1
            i++
            len++
        }else{
            if len > 0 {
                len = lsp[len-1]
            }else{
                lsp[i] = 0
                i++
            }
        }
    }

    return lsp
}

func kmpSolution(haystack, needle string) int {
    lsp := kmpPrefixTable(haystack, needle)
    fmt.Println("lsp", lsp)
    m, n := len(haystack), len(needle)

    i, j := 0, 0
    for i < m {
        if haystack[i] == needle[j] {
            i++
            j++
        }else{
            if j > 0 {
                j = lsp[j-1]
            }else{
                i++
            }
        }

        if j == n {
            return i-n
        }
    }
    return -1
}