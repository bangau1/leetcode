use std::vec::Vec;

pub struct DisjointSet {
  parents: Vec<i32>,
}

impl DisjointSet {
  fn new(n: u32) -> Self {
    return DisjointSet { 
      parents: vec![-1; n as usize], 
    }
  }

  fn find(&mut self, x: u32) -> u32 {
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

  fn union(&mut self, a: u32, b: u32) -> bool {
    if a == b {
      return false;
    }

    let (a_id, b_id) = (self.find(a) as i32, self.find(b) as i32);

    if a_id == b_id {
      return false
    }

    let (a_size, b_size) = (self.union_size(a_id as u32), self.union_size(b_id as u32));

    if b_size > a_size {
      self.parents[a_id as usize] = b_id as i32;
      self.parents[b_id as usize] = -((a_size+b_size) as i32);
    }else{
      self.parents[b_id as usize] = a_id as i32;
      self.parents[a_id as usize] = -((a_size+b_size) as i32);
    }
    return true
  }

  fn union_size(&mut self, x: u32) -> u32 {
    let x_id = self.find(x);
    return (-self.parents[x_id as usize]) as u32
  }

  fn total_set(&self) -> u32 {
    let mut result = 0;
    for &i in self.parents.iter() {
      if i < 0 {
        result += 1;
      }
    }
    return result as u32;
  }
}

#[test]
fn test_disjoint_set() {
  let mut ds = DisjointSet::new(3);
  assert_eq!(1, ds.find(1));
  assert_eq!(3, ds.total_set());

  ds.union(1, 2);
  assert_eq!(ds.find(1), ds.find(2));
  assert_eq!(2, ds.union_size(1));
  assert_ne!(ds.find(0), ds.find(1));
  assert_eq!(2, ds.total_set());
  
}