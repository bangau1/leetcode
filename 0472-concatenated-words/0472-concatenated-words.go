type trie struct {
    root *node
}

func newTrie(words []string) trie {
    tr := trie{
        root: &node{
            children: make([]*node, 26),
        },
    }
    for _, word := range words {
        tr.root.addWord(word)
    }
    return tr
}

type node struct {
    children []*node
    wordCount int
    prefixCount int
}

type parentToChild struct{
    parent *node
    childLetter int
}

func (this *node) addWord(word string) {
    if len(word) == 0 {
        this.wordCount++
        this.prefixCount++
        return
    }

    this.prefixCount++
    letter := int(word[0]-"a"[0])
    if this.children[letter] == nil {
        this.children[letter] = &node{
            children: make([]*node, 26),
        }
    }
    this.children[letter].addWord(word[1:])
}

func (this *trie) removeWord(word string) {
    var curr = this.root
    var parents = make([]parentToChild, 0)
    for i:=0;i<len(word);i++{
        letter := int(word[i]-"a"[0])
        if curr.children[letter] == nil {
            return
        }
        parents = append(parents, parentToChild{curr, letter})
        curr = curr.children[letter]
    }
    curr.wordCount--
    curr.prefixCount--
    // this below is optional, can be useful if the length of word is quite large, so we can remove unnecessary pointer
    for len(parents) > 0 {
        p := parents[len(parents)-1]
        parents = parents[0:len(parents)-1]
        
        if curr.prefixCount == 0 {
            p.parent.children[p.childLetter] = nil
        }
        p.parent.prefixCount--
        curr = p.parent
    }
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

    letter := int(word[0]-"a"[0])
    return this.children[letter].search(word[1:])
}

func wordBreakByTrie(tr trie, s string) bool {
    
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
            letter := int(s[start]-"a"[0])

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


func findAllConcatenatedWordsInADict(words []string) []string {
    tr := newTrie(words)
    res := make([]string, 0)
    sort.Slice(words, func(a, b int)bool {
        return len(words[a]) > len(words[b])
    })
    for i:=0;i<len(words)-1;i++{
        tr.removeWord(words[i])
        if wordBreakByTrie(tr, words[i]){
            res = append(res, words[i])
        }
        
    }
    return res
}