use std::vec::Vec;

fn next_greater_element(arr: &Vec<i32>, is_forward: bool) -> Vec<i32> {
  let mut nge: Vec<i32> = vec![-1;arr.len()];
  let mut stack: Vec<usize> = vec![];
  if is_forward {

    for (pos, &el) in arr.iter().enumerate() {
      while stack.len() > 0 && arr[*stack.last().unwrap()] < el {
        let last = stack.pop().unwrap();
        nge[last] = pos as i32;
      }
      stack.push(pos);
    }

  }else{
    for (pos, &el) in arr.iter().enumerate().rev() {
      while stack.len() > 0 && arr[*stack.last().unwrap()] < el {
        let last = stack.pop().unwrap();
        nge[last] = pos as i32;
      }
      stack.push(pos);
    }
  }

  while stack.len() > 0 {
    nge[stack.pop().unwrap()] = -1
  }

  return nge;
}

#[test]
fn test_nge() {
  let arr = vec![1,2,-1,10];
  let nge = next_greater_element(&arr, true);

  assert_eq!([1,3,3,-1], nge.as_slice());
}

#[test]
fn test_pge(){
  let arr = vec![1,2,-1,10];
  let nge = next_greater_element(&arr, false);

  assert_eq!([-1,-1,1,-1], nge.as_slice());
}

pub fn trap(height: Vec<i32>) -> i32 {
    let nge = next_greater_element(&height, true);
    let pge = next_greater_element(&height, false);

    let mut i = 0;
    let mut total_water = 0;
    while i < height.len() {
      let mut right = i;
      while nge[right] != -1 && pge[right] != -1 {
        right = nge[right] as usize;
      }

      let mut left = i;
      while pge[left] != -1 && nge[left] != -1 {
        left = pge[left] as usize;
      }

      let max = std::cmp::max(height[left], height[right]);
      let mut sub_total = 0;
      for pos in left..right+1 {
        sub_total += max - height[pos];
      }
      total_water += sub_total;

      i = right + 1;
    }
    return total_water

}