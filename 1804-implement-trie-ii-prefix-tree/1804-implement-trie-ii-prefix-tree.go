type Trie struct {
    root *Node
}

type Node struct {
    children []*Node
    wordCount int
    prefixCount int
}


func Constructor() Trie {
    return Trie{
        root: &Node{
            children: make([]*Node, 26),
            wordCount: 0,
        },
    }
    
}


func (this *Trie) Insert(word string)  {
    var curr = this.root
    curr.prefixCount++
    for i:=0;i<len(word);i++{
        pos := int(word[i] -"a"[0])
        if curr.children[pos] == nil {
            curr.children[pos] = &Node{
                children: make([]*Node, 26),
                wordCount:0,
            }
        }
        curr = curr.children[pos]
        curr.prefixCount++
    }
    curr.wordCount++
}


func (this *Trie) Search(word string) bool {
    var curr = this.root
    for i:=0;i<len(word);i++{
        pos := int(word[i] -"a"[0])
        if curr.children[pos] == nil {
           return false
        }
        curr = curr.children[pos]
    }
    return curr.wordCount > 0
}


func (this *Trie) StartsWith(prefix string) bool {
    var curr = this.root
    for i:=0;i<len(prefix);i++{
        pos := int(prefix[i] -"a"[0])
        if curr.children[pos] == nil {
           return false
        }
        curr = curr.children[pos]
    }
    return true
}


func (this *Trie) CountWordsEqualTo(word string) int {
    var curr = this.root
    for i:=0;i<len(word);i++{
        pos := int(word[i] -"a"[0])
        if curr.children[pos] == nil {
           return 0
        }
        curr = curr.children[pos]
    }
    return curr.wordCount
}


func (this *Trie) CountWordsStartingWith(prefix string) int {
    var curr = this.root
    for i:=0;i<len(prefix);i++{
        pos := int(prefix[i] -"a"[0])
        if curr.children[pos] == nil {
           return 0
        }
        curr = curr.children[pos]
    }
    return curr.prefixCount
}


func (this *Trie) Erase(word string)  {
    var curr = this.root
    var prev *Node
    stack := make([]*Node, 0)
    for i:=0;i<len(word);i++{
        pos := int(word[i] -"a"[0])
        if curr.children[pos] == nil {
            return
        }
        prev = curr
        stack = append(stack, prev)
        curr = curr.children[pos]
    }

    if curr.wordCount == 0 {
        return
    }
    curr.prefixCount--
    curr.wordCount--
    
    var parent *Node
    i := len(word)-1
    for len(stack) > 0 {
        parent = stack[len(stack)-1]
        stack = stack[0:len(stack)-1]

        parent.prefixCount-- // we decrease the prefix along the path
        if curr.prefixCount == 0 {
            pos := int(word[i]-"a"[0])
            parent.children[pos] = nil
        }
        curr = parent
        i--
    }
    
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */