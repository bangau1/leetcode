type trie struct {
    root *node
}

func newTrie(words []string) trie {
    tr := trie{
        root: &node{
            children: make(map[rune]*node),
        },
    }
    for _, word := range words {
        tr.root.addWord(word)
    }
    return tr
}

type node struct {
    children map[rune]*node
    wordCount int
    prefixCount int
}

func (this *node) addWord(word string) {
    if len(word) == 0 {
        this.wordCount++
        this.prefixCount++
        return
    }

    this.prefixCount++
    letter := rune(word[0])
    if this.children[letter] == nil {
        this.children[letter] = &node{
            children: make(map[rune]*node),
        }
    }
    this.children[letter].addWord(word[1:])
}

func (this *node) search(word string) bool {
    if this == nil && len(word) > 0 {
        return false
    }

    if len(word) == 0 {
        if this != nil {
            return this.wordCount > 0
        }else{
            return false
        }
    }

    letter := rune(word[0])
    return this.children[letter].search(word[1:])
}
 
func (this *node) searchLongestMatch(word string) int {
    var curr = this
    maxIdx := -1
    
    for i:=0;i<len(word);i++{
        letter := rune(word[i])
        if curr.children[letter] == nil {
            break
        }
        curr = curr.children[letter]
        if curr.wordCount > 0 {
            maxIdx = i
        }
    }
    return maxIdx
}

func (this *trie) isSegmentedWord(word string) bool {
    if len(word) == 0 {
        return true
    }
    var curr = this.root
    res := false
    stack := make([]string, 0)
    for i:=0;i<len(word);i++{
        letter := rune(word[i])
       
        if curr.children[letter] == nil {
            break
        }

        curr = curr.children[letter]
        if curr.wordCount > 0 {
            stack = append(stack, word[i+1:])
        }
    }
    if len(stack) == 0 {
        return false
    }
    for len(stack) > 0 {
        item := stack[len(stack)-1]
        fmt.Println(item)
        stack = stack[0:len(stack)-1]
        if this.isSegmentedWord(item) {
            return true
        }
    }
    return res
}

func wordBreakBySimpleDP(s string, wordDict []string) bool {
    // let dp[i] = means we can build up to i index (exclusive) using word in the dictionary
    // dp[i] = for each w[a] (w[a] in wordDict), dp[i-len(w[a])] && s[i-len(w[a]:i)] == word
    dp := make([]bool, len(s)+1)
    dp[0] = true // empty string is true
    for i:=0;i<len(s);i++{
        for _, w := range wordDict {
            wn := len(w) // 4
            prevIdx := i - wn +1 // i = 3  prevIdx = 2
            if prevIdx < 0 {
                continue
            }
            //s[3:5] = s[3],s[4]
            if len(w) != len(s[prevIdx:i+1]){
                panic("unexpected")
            }
            if dp[prevIdx] && s[prevIdx:i+1] == w {
                dp[i+1] = true
                break
            }
        }
    }
    // fmt.Println(dp[1:])
    return dp[len(s)]
}

func wordBreak(s string, wordDict []string) bool {
    return wordBreakByTrie(s, wordDict)
}

func wordBreakByTrie(s string, wordDict []string) bool {
    tr := newTrie(wordDict)

    // let dp[i] = means we can build up to i index (exclusive) using word in the dictionary
    // dp[i] = for each w[a] (w[a] in wordDict), dp[i-len(w[a])] && s[i-len(w[a]:i)] == word
    dp := make([]bool, len(s)+1)
    dp[0] = true // empty string is true
    for i:=0;i<len(s);i++{
        if !dp[i] {
            continue
        }

        // from starting index i, we will traverse the trie and mark the dp[j] to true if the word is there
        var curr = tr.root
        for start:=i; start < len(s); start++{
            letter := rune(s[start])

            if curr.children[letter] == nil {
                break
            }
            curr = curr.children[letter]
            if curr.wordCount > 0 {
                dp[start+1] = true
            }
        }
    }
    // fmt.Println(dp[1:])
    return dp[len(s)]
}

