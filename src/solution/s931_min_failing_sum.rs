struct Solution {}
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

#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn test_case_1() {
        assert_eq!(
            13,
            Solution::min_falling_path_sum(vec![vec![2, 1, 3], vec![6, 5, 4], vec![7, 8, 9]])
        );
    }

    fn test_case_2() {
        assert_eq!(
            -59,
            Solution::min_falling_path_sum(vec![vec![-19, 57], vec![-40, -5]])
        );
    }
}
