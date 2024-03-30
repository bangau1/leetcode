
func countVowelSubstrings(word string) int {
    return countAtMostKVowels(word, 5) - countAtMostKVowels(word, 4)
}

func countAtMostKVowels(word string, k int) int {
    counter := make(map[byte]int)
    var total int
    var l int
    for r := 0;r<len(word);r++{
        if !isVowel(word[r]) {
            l = r + 1
            counter = make(map[byte]int)
            continue
        }
        counter[word[r]]++
        for l<=r && len(counter) > k {
            counter[word[l]]--
            if counter[word[l]] == 0 {
                delete(counter, word[l])
            }
            l++
        }
        total += r-l+1
    }
    return total
}

func isVowel(char byte) bool {
    return char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o'
}