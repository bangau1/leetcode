type Wheels struct {
    data [4]byte
}

func (w *Wheels) String() string {
    res := make([]byte, 0)
    res[0] = w.data[0]
    
}

type Vertex struct {
    w Wheels
    c int
}

func NewWheels(str string) Wheels {
    return Wheels {
        data: [4]byte{str[0], str[1], str[2], str[3]},
    }
}

func (w *Wheels) getNextSteps() []Wheels {
    
    res := make([]Wheels, 0)
    for i:=0;i<4;i++{
        currDigit := int(w.data[i] - '0')
        currDigit = (currDigit + 1) % 10
        newWheel := *w
        newWheel.data[i] = byte(currDigit) + '0'
        res = append(res, newWheel)

        currDigit = int(w.data[i] - '0')
        currDigit = currDigit -1
        if currDigit < 0 {
            currDigit += 10
        }

        newWheel = *w
        newWheel.data[i] = byte(currDigit) + '0'
        res = append(res, newWheel)
    }
    return res
}


func openLock(deadends []string, target string) int {
    ignore := make(map[Wheels]bool)

    for _, d := range deadends {
        w := NewWheels(d)
        ignore[w] = true
    }

    visited := make(map[Wheels]bool)

    q := make([]Vertex, 0)
    q = append(q, Vertex{NewWheels("0000"), 0})
    
    t := NewWheels(target)
    
    var node Vertex
    for len(q) > 0 {
        node = q[0]
        q = q[1:]

        if visited[node.w] {
            continue
        } 

        visited[node.w] = true
        if node.w == t {
            return node.c
        }

        for _, next := range node.w.getNextSteps() {
            if !visited[next] && !ignore[next] {
                fmt.Println(string(next.data))
                q = append(q, Vertex{next, node.c+1})
            }
        }
    }
    return -1
}