func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var counter [26]int
    var zeroes [26]int
    for _, char := range s {
        counter[int(char-'a')]++
    }

    for _, char := range t {
        counter[int(char-'a')]--
        if counter[int(char)-'a'] < 0 {
            return false
        }
    }
    return counter == zeroes
}