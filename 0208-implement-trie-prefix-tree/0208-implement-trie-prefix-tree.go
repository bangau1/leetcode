type Trie struct {
    root *Node
}

type Node struct {
    children []*Node
    wordCount int
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
    for i:=0;i<len(word);i++{
        pos := int(word[i] -"a"[0])
        if curr.children[pos] == nil {
            curr.children[pos] = &Node{
                children: make([]*Node, 26),
                wordCount:0,
            }
        }
        curr = curr.children[pos]
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


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */