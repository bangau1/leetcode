use std::{cmp::Ordering, collections::BinaryHeap};

pub fn djikstra(src: usize, adj_list: Vec<Vec<Vertex>>) -> Vec<i32> {
    let n = adj_list.len();

    let mut dist = vec![i32::MAX; n];
    dist[src] = 0;

    let mut visited = vec![false; n];

    let mut min_heap = BinaryHeap::new();
    min_heap.push(Vertex { node: src, cost: 0 });

    while min_heap.len() > 0 {
        let u = min_heap.pop().unwrap();
        if visited[u.node] {
            continue;
        }

        visited[u.node] = true;

        for &edge in adj_list[u.node].iter() {
            if !visited[edge.node] && dist[edge.node] > (dist[u.node] + edge.cost) {
                dist[edge.node] = dist[u.node] + edge.cost;
                min_heap.push(Vertex {
                    node: edge.node,
                    cost: dist[edge.node],
                })
            }
        }
    }

    return dist;
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Vertex {
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
            return other.node.cmp(&self.node);
        }
        return other.cost.cmp(&self.cost);
    }
}
