func NewLetterCount(str string) [26]byte {
    counter := [26]byte{}

    for i:=0;i<len(str);i++{
        counter[int(str[i]-'a')]++
    }
    return counter
}

func groupAnagrams(strs []string) [][]string {
    group := make(map[[26]byte][]string)
    var lc [26]byte
    for _, str := range strs {
        lc = NewLetterCount(str)
        group[lc] = append(group[lc], str)
    }
    res := make([][]string, len(group))
    i := 0
    for _, list := range group {
        res[i] = list
        i++
    }

    return res
}