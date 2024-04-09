use std::collections::BinaryHeap;
use std::cmp::Ordering;

impl Solution {
    
    pub fn network_delay_time(times: Vec<Vec<i32>>, n: i32, k: i32) -> i32 {
        // calculate the shortest path from the source to all nodes (can use djikstra)
        // then calculate the max of dists
        let n = n as usize;
        let k = k as usize;

        let mut adjList: Vec<Vec<Vertex>> = vec![Vec::<Vertex>::new();n as usize];
        for t in times.iter() {
            let (src, dst, cost) = ((t[0]-1) as usize, (t[1]-1) as usize, t[2]);
            adjList[src].push(Vertex{node: dst, cost: cost});
        }
        let dists = Solution::djikstra(k-1, adjList);
        

        let &max_time = dists.iter().max().unwrap();
        if max_time == i32::MAX {
            return -1
        }
        return max_time;
    }

    pub fn djikstra(src: usize, adjList: Vec<Vec<Vertex>>) -> Vec<i32> {
        let n = adjList.len();

        let mut dist = vec![i32::MAX;n];
        dist[src] = 0;

        let mut visited = vec![false;n];

        let mut minHeap = BinaryHeap::new();
        minHeap.push(Vertex{node: src, cost: 0});

        while minHeap.len() > 0 {
            let u = minHeap.pop().unwrap();
            if visited[u.node] {
                continue;
            }

            visited[u.node] = true;

            for &edge in adjList[u.node].iter() {
                if !visited[edge.node] && dist[edge.node] > (dist[u.node] + edge.cost) {
                    dist[edge.node] = dist[u.node] + edge.cost;
                    minHeap.push(Vertex{
                        node: edge.node,
                        cost: dist[edge.node],
                    })
                }
            }

        }

        return dist;
    }

}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct Vertex {
    node: usize,
    cost: i32,
}

impl PartialOrd for Vertex {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        return Some(self.cmp(other));
    }
}

impl Ord for Vertex {
    fn cmp(&self, other: &Self) -> Ordering {
        // we want to prioritize the minimum cost
        // since binaryHeap in rust is maxHeap, we should reverse it
        if self.cost == other.cost {
            return other.node.cmp(&self.node)
        }
        return other.cost.cmp(&self.cost);
    }
}