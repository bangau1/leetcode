type WordFilter struct {
    prefixRoot *node
    suffixRoot *node
}

type node struct {
    children map[rune]*node
    prefixLoc []int
    val string
}

func (this *node) searchPrefix(word string) []int{
    if this == nil {
        return nil
    }

    if len(word) == 0 {
        return this.prefixLoc
    }
    letter := rune(word[0])
    return this.children[letter].searchPrefix(word[1:])
}

func (this *node) searchSuffix(word string) []int{
    if this == nil {
        return nil
    }

    if len(word) == 0 {
        return this.prefixLoc
    }
    letter := rune(word[len(word)-1])
    return this.children[letter].searchSuffix(word[0:len(word)-1])
}

func (this *node) addPrefix(word string, index int) {
    if this == nil || len(word) == 0 {
        return
    }
    letter := rune(word[0])
    if this.children[letter] == nil {
        this.children[letter] = &node{
            children: make(map[rune]*node),
            val: string(word[0]),
        }
    }
    this.children[letter].prefixLoc = append(this.children[letter].prefixLoc, index)
    this.children[letter].addPrefix(word[1:], index)
}

func (this *node) addSuffix(word string, index int) {
    if this == nil || len(word) == 0 {
        return
    }
    
    letter := rune(word[len(word)-1])
    if this.children[letter] == nil {
        this.children[letter] = &node{
            children: make(map[rune]*node),
            val: string(word[0]),
        }
    }
    this.children[letter].prefixLoc = append(this.children[letter].prefixLoc, index)
    this.children[letter].addSuffix(word[0:len(word)-1], index)
}

type wordIdx struct{
    word string
    idx int
}

func Constructor(words []string) WordFilter {
    wordMap := make(map[string]int)
    for i, word := range words {
        wordMap[word] = i
    }

    wordSet := make([]wordIdx, 0)
    for word, idx := range wordMap {
        wordSet = append(wordSet, wordIdx{word, idx})
    }

    sort.Slice(wordSet, func(a, b int) bool {
        return wordSet[a].idx < wordSet[b].idx
    })

    wf := &WordFilter{
        prefixRoot: &node{
            children: make(map[rune]*node),
        },
        suffixRoot: &node{
            children: make(map[rune]*node),
        },
    }   

    for _, wordIdx := range wordSet {
        word, i := wordIdx.word, wordIdx.idx
        wf.prefixRoot.addPrefix(word, i)
        wf.suffixRoot.addSuffix(word, i)
    }
    return *wf
}


func (this *WordFilter) F(pref string, suff string) int {
    prefixLoc := this.prefixRoot.searchPrefix(pref)
    suffixLoc := this.suffixRoot.searchSuffix(suff)

    if len(prefixLoc) == 0 || len(suffixLoc) == 0 {
        return -1
    }

    // if no overlapping
    pMin, pMax := prefixLoc[0], prefixLoc[len(prefixLoc)-1]
    sMin, sMax := prefixLoc[0], prefixLoc[len(prefixLoc)-1]

    if pMin <= sMin {
        if pMax < sMin {
            return -1
        }
    }

    if sMin <= pMin {
        if sMax < pMin {
            return -1
        }
    }

    // fmt.Println(prefixLoc)
    // fmt.Println(suffixLoc)
    
    a, b := prefixLoc, suffixLoc
    if len(a) > len(b) {
        a, b = b, a
    }
    flag := make(map[int]bool)
    for _, it := range a {
        flag[it] = true
    }

    for i:=len(b)-1;i>=0;i--{
        if flag[b[i]] {
            return b[i]
        }
    }
    return -1
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */