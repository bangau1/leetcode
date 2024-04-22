use std::collections::{HashSet, VecDeque};

struct Solution {}
impl Solution {
    pub fn open_lock(deadends: Vec<String>, target: String) -> i32 {
        let mut ignore_map: HashSet<Wheel> = HashSet::new();
        for deadend in deadends {
            ignore_map.insert(Wheel::new(&deadend));
        }

        let target = Wheel::new(&target);
        let source = Wheel::new("0000");
        if ignore_map.contains(&target) || ignore_map.contains(&source) {
            return -1;
        }

        if target == source {
            return 0;
        }

        let mut queue: VecDeque<Vertex> = VecDeque::new();
        queue.push_back(Vertex {
            wheel: source,
            cost: 0,
        });
        let mut visited: HashSet<Wheel> = HashSet::new();

        while queue.len() > 0 {
            let node = queue.pop_front().unwrap();

            if visited.contains(&node.wheel) {
                continue;
            }

            visited.insert(node.wheel.clone());
            if node.wheel == target {
                return node.cost;
            }

            for next in node.wheel.next_steps() {
                if !visited.contains(&next) && !ignore_map.contains(&next) {
                    queue.push_back(Vertex {
                        wheel: next,
                        cost: node.cost + 1,
                    })
                }
            }
        }

        return -1;
    }
}

#[derive(Copy, Clone, PartialEq, Eq)]
struct Vertex {
    wheel: Wheel,
    cost: i32,
}

#[derive(Copy, Clone, PartialEq, Eq, Hash)]
struct Wheel([i8; 4]);
impl Wheel {
    fn new(str: &str) -> Wheel {
        let str_byte = str.as_bytes();
        return Wheel([
            (str_byte[0] - '0' as u8) as i8,
            (str_byte[1] - '0' as u8) as i8,
            (str_byte[2] - '0' as u8) as i8,
            (str_byte[3] - '0' as u8) as i8,
        ]);
    }

    fn next_steps(&self) -> Vec<Wheel> {
        let mut res = vec![];
        for i in 0..4 as usize {
            let mut curr = self.clone();
            curr.0[i] = (curr.0[i] + 1 as i8) % 10;

            res.push(curr);

            let mut curr = self.clone();
            curr.0[i] = (curr.0[i] - 1 as i8);
            if curr.0[i] < 0 {
                curr.0[i] += 10;
            }

            res.push(curr);
        }
        return res;
    }
}
