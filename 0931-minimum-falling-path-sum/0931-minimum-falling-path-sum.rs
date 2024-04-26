use std::cmp::{max, min};

impl Solution {
    pub fn min_falling_path_sum(matrix: Vec<Vec<i32>>) -> i32 {
        if matrix.len() == 1 {
            return *matrix[0].iter().min().unwrap();
        }
        let mut prev_sums = matrix[0].clone();
        for r in 1..matrix.len() {
            let mut curr_sums = vec![0; prev_sums.len()];
            for (c, &grid_val) in matrix[r].iter().enumerate() {
                let left = max(0, c as i32 - 1) as usize;
                let right = min(curr_sums.len() - 1, c + 1);

                curr_sums[c] = prev_sums[left..right + 1].iter().min().unwrap() + grid_val;
            }
            prev_sums = curr_sums;
        }
        return prev_sums.into_iter().min().unwrap();
    }
}