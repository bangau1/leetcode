func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    wordMap := make(map[string]int)
    for i, word := range wordList {
        wordMap[word]=i
    }

    if _, ok:=wordMap[endWord]; !ok {
        return nil
    }

    n := len(wordList)
    if _, ok:=wordMap[beginWord]; !ok {
        wordMap[beginWord] = n
        wordList = append(wordList, beginWord)
        n += 1
    }
    visited := make([]bool, n)
    prevNode := make([]int, n)
    for i:=0;i<n;i++{
        prevNode[i] = -1
    }

    bfs := make([]vertex, 0)
    bfs = append(bfs, vertex{
        wordMap[beginWord], wordMap[beginWord], 1,
    })
    target := wordMap[endWord]

    dist := math.MaxInt
    dists := make([]int, n)
    for i:=0;i<n;i++{
        dists[i] = math.MaxInt
    }
    dists[wordMap[beginWord]] = 0
    for len(bfs) > 0 {
        cur := bfs[0]
        bfs = bfs[1:]

        if visited[cur.node] {
            continue
        }
        visited[cur.node] = true
        prevNode[cur.node] = cur.prev
        dists[cur.node] = min(dists[cur.node], cur.length - 1)
        if cur.node == target {
            dist = cur.length
            continue
        }
        if cur.length + 1 > dist {
            continue
        }
        // now iterate the next word
        for _, nextIdx := range generateNextWord(wordMap, wordList[cur.node]) {
            if !visited[nextIdx] && cur.length + 1 <= dist {
                bfs = append(bfs, vertex{
                    cur.node, nextIdx, cur.length + 1,
                })
            }
        }

    }

    if dist == math.MaxInt {
        return nil
    }
    src := wordMap[beginWord]
    // fmt.Println("dist", dist)
    // fmt.Println("dists", dists)
    var res [][]string
    visited = make([]bool, n)
    var dfs func(int, int, []int)
    dfs = func(currNode int, level int, path []int){
        if level > dist {
            return
        }
        if visited[currNode] {
            return
        }
        visited[currNode] = true
        if currNode == src {
            tmp := make([]string, len(path))
            for i:=0;i<len(path);i++{
                tmp[i] = wordList[path[ len(path)-1-i]]
            }
            res = append(res, tmp)
            return
        }

        if level + 1 > dist {
            return
        }

        // now iterate the next word
        for _, nextIdx := range generateNextWord(wordMap, wordList[currNode]) {
            if !visited[nextIdx] && dists[nextIdx] != math.MaxInt && level + 1 + dists[nextIdx] <= dist {
                path = append(path, nextIdx)
                dfs(nextIdx, level + 1, path)
                path = path[:len(path)-1]
                visited[nextIdx] = false
            }
        }

    }
    dfs(target, 1, []int{target})

    return res
}



func getPath(wordList []string, prevNode []int, curr int) []string {
    res := make([]string, 0)
    for true {
        res = append(res, wordList[curr])
        if curr == prevNode[curr] {
            break
        }
        curr = prevNode[curr]
    }
    slices.Reverse(res)
    return res
}

// assuming the word length is <= 5, which is small enough
func generateNextWord(wordIdx map[string]int, word string) []int{
    var res []int
    n := len(word)
    for i:=0;i<n;i++{
        data := []byte(word)
        char := data[i]
        for c:=byte(0);c<byte(26);c++{
            data[i] = byte('a') + c
            if char == data[i] {
                continue
            }
            if idx, ok:=wordIdx[string(data)]; ok {
                res = append(res, idx)
            }
        }
    }
    return res
}

type vertex struct {
    prev int
    node int
    length int
}