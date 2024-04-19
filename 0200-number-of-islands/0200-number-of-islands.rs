impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let size = grid.len() * grid[0].len();
        let mut ds = DisjointSet::new(size as u32);

        for r in 0..grid.len() {
            for c in 0..grid[r].len() {
                if grid[r][c] == '1' {
                    let curr_cell_id = to_cell_id(r, c, grid[0].len());
                    if r >= 1 && grid[r - 1][c] == '1' {
                        ds.union(curr_cell_id, to_cell_id(r - 1, c, grid[0].len()));
                    }
                    if c >= 1 && grid[r][c - 1] == '1' {
                        ds.union(curr_cell_id, to_cell_id(r, c - 1, grid[0].len()));
                    }
                }else {
                    let curr_cell_id = to_cell_id(r, c, grid[0].len());
                    if r >= 1 && grid[r - 1][c] != '1' {
                        ds.union(curr_cell_id, to_cell_id(r - 1, c, grid[0].len()));
                    }
                    if c >= 1 && grid[r][c - 1] != '1' {
                        ds.union(curr_cell_id, to_cell_id(r, c - 1, grid[0].len()));
                    }
                }
            }
        }
        let mut total = 0;
        let union_ids = ds.get_all_union_ids();
        for &id in union_ids.iter() {
            let (r, c) = cell_id_to_row_col(id as u32, grid[0].len());
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

pub struct DisjointSet {
    parents: Vec<i32>,
}

impl DisjointSet {
    pub fn new(n: u32) -> Self {
        return DisjointSet {
            parents: vec![-1; n as usize],
        };
    }

    pub fn find(&mut self, x: u32) -> u32 {
        let mut x = x as usize;
        if self.parents[x] < 0 {
            return x as _;
        }

        let mut path = vec![];
        while self.parents[x] >= 0 {
            path.push(x);
            x = self.parents[x] as usize;
        }

        for &p in path.iter() {
            self.parents[p] = x as i32;
        }
        return x as _;
    }

    pub fn union(&mut self, a: u32, b: u32) -> bool {
        if a == b {
            return false;
        }

        let (a_id, b_id) = (self.find(a) as i32, self.find(b) as i32);

        if a_id == b_id {
            return false;
        }

        let (a_size, b_size) = (self.union_size(a_id as u32), self.union_size(b_id as u32));

        if b_size > a_size {
            self.parents[a_id as usize] = b_id as i32;
            self.parents[b_id as usize] = -((a_size + b_size) as i32);
        } else {
            self.parents[b_id as usize] = a_id as i32;
            self.parents[a_id as usize] = -((a_size + b_size) as i32);
        }
        return true;
    }

    pub fn union_size(&mut self, x: u32) -> u32 {
        let x_id = self.find(x);
        return (-self.parents[x_id as usize]) as u32;
    }

    pub fn total_set(&self) -> u32 {
        let mut result = 0;
        for &i in self.parents.iter() {
            if i < 0 {
                result += 1;
            }
        }
        return result as u32;
    }

    pub fn get_all_union_ids(&self) -> Vec<usize> {
        let mut union_ids = vec![];
        for (i, &val) in self.parents.iter().enumerate() {
            if val < 0 {
                union_ids.push(i);
            }
        }
        return union_ids;
    }
}