struct Solution {}

use std::collections::BinaryHeap;

#[derive(PartialEq, Eq, Copy, Clone, Debug)]
struct Record {
    value: i32,
    index: usize,
}

impl Ord for Record {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        if self.value == other.value {
            return self.index.cmp(&other.index);
        }
        return self.value.cmp(&other.value);
    }
}

impl PartialOrd for Record {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        return Some(self.cmp(other));
    }
}

impl Solution {
    pub fn min_falling_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        if grid.len() == 1 {
            return *grid[0].iter().min().unwrap();
        }

        // maintain max heap with size 2, to contains 2 minimum data in the heap
        let mut prev_sum: BinaryHeap<Record> = BinaryHeap::new();
        for (pos, &grid_val) in grid[0].iter().enumerate() {
            if prev_sum.len() < 2 {
                prev_sum.push(Record {
                    value: grid_val,
                    index: pos,
                });
            } else if grid_val < prev_sum.peek().unwrap().value {
                prev_sum.pop();
                prev_sum.push(Record {
                    value: grid_val,
                    index: pos,
                });
            }
        }

        for r in 1..grid.len() {
            let mut curr_sum: BinaryHeap<Record> = BinaryHeap::new();
            let prev_mins: Vec<Record> = Vec::from_iter(prev_sum.into_iter());
            println!("prev_mins = {:?}", prev_mins);

            for (c, &grid_val) in grid[r].iter().enumerate() {
                let mut grid_sum = Record {
                    value: grid_val,
                    index: c,
                };
                if grid_sum.index != prev_mins[1].index {
                    grid_sum.value += prev_mins[1].value;
                } else {
                    grid_sum.value += prev_mins[0].value;
                }

                if curr_sum.len() < 2 {
                    curr_sum.push(grid_sum);
                } else if grid_sum.value < curr_sum.peek().unwrap().value {
                    curr_sum.pop();
                    curr_sum.push(grid_sum);
                }
            }
            prev_sum = curr_sum;
        }

        println!("prev_mins = {:?}", prev_sum);
        return prev_sum.iter().min().unwrap().value;
    }

    pub fn min_falling_path_sum_n3(grid: Vec<Vec<i32>>) -> i32 {
        if grid.len() == 1 {
            return *grid[0].iter().min().unwrap();
        }

        let mut prev_sum = vec![0; grid[0].len()];
        for r in 0..grid.len() {
            let mut curr_sum = vec![0; prev_sum.len()];

            for (c, &grid_val) in grid[r].iter().enumerate() {
                let mut curr_min = i32::MAX;
                for (prev, &prev_sum_val) in prev_sum.iter().enumerate() {
                    if prev != c && curr_min > grid_val + prev_sum_val {
                        curr_min = grid_val + prev_sum_val;
                    }
                }
                curr_sum[c] = curr_min;
            }
            prev_sum = curr_sum;
        }
        return *prev_sum.iter().min().unwrap();
    }
}

#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn test_cases() {
        assert_eq!(
            13,
            Solution::min_falling_path_sum(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]])
        );
        assert_eq!(7, Solution::min_falling_path_sum(vec![vec![7]]));
    }
}
