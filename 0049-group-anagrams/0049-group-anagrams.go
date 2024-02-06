func NewLetterCount(str string) string {
    counter := make([]int, 26)
    for i:=0;i<len(str);i++{
        letter := int(str[i]-'a')
        counter[letter]++
    }
    res := ""
    for i:=0;i<26;i++{
        if counter[i] != 0 {
            res = res + fmt.Sprintf("%s%d", byte(i)+'a', counter[i])
        }
    }
    return res
}

func groupAnagrams(strs []string) [][]string {
    hashmap := make(map[string][]string)
    for _, str := range strs {
        lc := NewLetterCount(str)
        hashmap[lc] = append(hashmap[lc], str)
    }
    res := make([][]string, 0)
    for _, list := range hashmap {
        res = append(res, list)
    }

    return res
}