use std::collections::HashSet;

struct Solution {}
impl Solution {
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
    pub fn find_min_height_trees(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        if n == 1 {
            return vec![0];
        }
        if n == 2 {
            return vec![0, 1];
        }

        let mut adj_list: Vec<HashSet<i32>> = vec![HashSet::new(); n as usize];

        for edge in edges {
            let (a, b) = (edge[0] as usize, edge[1] as usize);
            adj_list[a].insert(b as i32);
            adj_list[b].insert(a as i32);
        }

        let mut queue: Vec<i32> = vec![];
        for (i, edge_set) in adj_list.iter().enumerate() {
            if edge_set.len() == 1 {
                queue.push(i as i32);
            }
        }

        let min_remove_count = n - 2;
        let mut remove_count = 0;

        while remove_count < min_remove_count {
            let mut new_queue = vec![];
            for node in queue {
                let node = node as usize;
                let &next = adj_list[node].iter().next().unwrap();
                adj_list[next as usize].remove(&(node as i32));
                if adj_list[next as usize].len() == 1 {
                    new_queue.push(next);
                }
                remove_count += 1;
            }
            queue = new_queue;
        }
        return queue;
    }
}
