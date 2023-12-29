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


func wordBreak(s string, wordDict []string) []string {
    tr := newTrie(wordDict)

    // let dp[i] = means we can build up to i index (exclusive) using word in the dictionary
    // dp[i] = for each w[a] (w[a] in wordDict), dp[i-len(w[a])] && s[i-len(w[a]:i)] == word
    dp := make([]bool, len(s)+1)
    dp[0] = true // empty string is true

    adjList := make([][]int, len(s)+1)
    for i:=0;i<len(s);i++{
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
                adjList[i] = append(adjList[i], start+1)
            }
        }
    }
    // fmt.Println(dp[1:])
    if !dp[len(s)] {
        return nil
    }

    // fmt.Println(adjList)

    pq := make([]vertex, 0)
    pq = append(pq, vertex{0, ""})
    res := make([]string, 0)
    
    for len(pq) > 0 {
        node := pq[0]
        pq = pq[1:]

        if node.end == len(s) {
            res = append(res, node.segment)
            continue
        }

        for _, next := range adjList[node.end] {
            var newSegment string 
            if len(node.segment) > 0 {
                newSegment = node.segment + " " + s[node.end:next]
            }else{
                newSegment = s[node.end:next]
            }
            pq = append(pq, vertex{next, newSegment})
        }

    }

    return res
}

type vertex struct {
    end int
    segment string
}