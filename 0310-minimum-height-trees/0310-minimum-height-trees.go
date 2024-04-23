var empty struct{};

type node struct {
    label int
    edges map[int]struct{}
}

func findMinHeightTrees(n int, edges [][]int) []int {
/*
Observations:
- it seems that the MHT root's candidate is the one who has edgeCount than the rest
- but the example/case 2 is contradictive with that
- so I guess it's just to process the node with less edgeCount first, then proceeed until we found the root

In other word: topology sorting.
- Process with node that has less edgeCount, remove them one by one
- proceed with the next set of nodes
- the root will be found when there is no edge anymore (or a tie between 2 nodes, like case 2)
*/

    if n == 1 {
        return []int{0}
    }

    if n == 2 {
        return []int{0, 1}
    }

    nodes := make([]node, n)
    for i:=0;i<n;i++{
        nodes[i] = node {
            label: i,
            edges: make(map[int]struct{}),
        }
    }
    var a, b int
    for _, edge := range edges {
        a, b = edge[0], edge[1]
        nodes[a].edges[b] = empty
        nodes[b].edges[a] = empty
    }

    queue := make([]int, 0)

    for _, nodeC := range nodes {
        if len(nodeC.edges) == 1 {
            queue = append(queue, nodeC.label)
        }
    }
    var curr int
    shouldRemoveMinimum := n-2
    removeCount := 0

    for removeCount < shouldRemoveMinimum {
        newQueue := make([]int, 0)
        
        for len(queue) > 0 {
            curr = queue[0]
            queue = queue[1:]

            for next, _ := range nodes[curr].edges {
                removeCount++
                delete(nodes[next].edges, curr)

                if len(nodes[next].edges) == 1 {
                    newQueue = append(newQueue, next)
                }
            }
        }
        queue = newQueue
    }

    return queue
}