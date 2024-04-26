
impl Solution {
    pub fn min_falling_path_sum(grid: Vec<Vec<i32>>) -> i32 {
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
