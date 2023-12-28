type WordDictionary struct {
    root *node
}

type node struct {
    children map[rune]*node
    wordCount int
}


func Constructor() WordDictionary {
    return WordDictionary{
        root: &node{
            children: make(map[rune]*node),
            wordCount: 0,
        },
    }
}


func (this *WordDictionary) AddWord(word string)  {
    curr := this.root

    for _, letter := range word {
        
        if curr.children[letter] == nil{
            curr.children[letter] = &node{
                children: make(map[rune]*node),
                wordCount: 0,
            }
        }
        curr = curr.children[letter]
    }
    curr.wordCount++
}


func (this *WordDictionary) Search(word string) bool {
    return this.root.search(word)
}

func (this *node) search(word string) bool {
    if this == nil && len(word) > 0 {
        return false
    }

    if len(word) == 0 {
        return this != nil && this.wordCount > 0
    }
    if word[0] == '.' {
        for _, child := range this.children{
            if child.search(word[1:]) {
                return true
            }
        }

        return false
    }
    return this.children[rune(word[0])].search(word[1:])
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */