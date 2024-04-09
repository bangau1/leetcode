use std::vec::Vec;
impl Solution {
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

    pub fn trap(height: Vec<i32>) -> i32 {
        let nge = Solution::next_greater_element(&height, true);
        let pge = Solution::next_greater_element(&height, false);
      
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

            let min = std::cmp::min(height[left], height[right]);
            if left != i && right != i {
                for pos in left+1..right {
                    total_water += min - height[pos];
                }
                i = right + 1;
            } else{
                i+=1;
            }

        }
        return total_water

    }
}