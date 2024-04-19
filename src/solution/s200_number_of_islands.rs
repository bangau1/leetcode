use crate::graph::disjoint_set::DisjointSet;
pub struct Solution {}
impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let size = grid.len() * grid[0].len();
        let mut ds = DisjointSet::new(size as u32);
        let col_size = grid[0].len();
        for r in 0..grid.len() {
            for c in 0..grid[r].len() {
                if grid[r][c] == '1' {
                    let curr_cell_id = to_cell_id(r, c, col_size);
                    if r >= 1 && grid[r - 1][c] == '1' {
                        ds.union(curr_cell_id, to_cell_id(r - 1, c, col_size));
                    }
                    if c >= 1 && grid[r][c - 1] == '1' {
                        ds.union(curr_cell_id, to_cell_id(r, c - 1, col_size));
                    }
                } else {
                    let curr_cell_id = to_cell_id(r, c, grid[0].len());
                    if r >= 1 && grid[r - 1][c] != '1' {
                        ds.union(curr_cell_id, to_cell_id(r - 1, c, col_size));
                    }
                    if c >= 1 && grid[r][c - 1] != '1' {
                        ds.union(curr_cell_id, to_cell_id(r, c - 1, col_size));
                    }
                }
            }
        }
        let mut total = 0;
        let union_ids = ds.get_all_union_ids();
        for &id in union_ids.iter() {
            let (r, c) = cell_id_to_row_col(id as u32, col_size);
            if grid[r][c] == '1' {
                total += 1;
            }
        }
        return total;
    }
}

fn to_cell_id(row: usize, col: usize, col_size: usize) -> u32 {
    return (row * col_size + col) as u32;
}

fn cell_id_to_row_col(cell_id: u32, col_size: usize) -> (usize, usize) {
    let col_size = col_size as u32;
    let row = cell_id / col_size;
    let col = cell_id % col_size;

    return (row as usize, col as usize);
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test() {
        assert_eq!(
            1,
            Solution::num_islands(vec![
                vec!['1', '1', '1', '1', '0'],
                vec!['1', '1', '0', '1', '0'],
                vec!['1', '1', '0', '0', '0'],
                vec!['0', '0', '0', '0', '0']
            ])
        );

        assert_eq!(
            3,
            Solution::num_islands(vec![
                vec!['1', '1', '0', '0', '0'],
                vec!['1', '1', '0', '0', '0'],
                vec!['0', '0', '1', '0', '0'],
                vec!['0', '0', '0', '1', '1'],
            ])
        );
    }
}
