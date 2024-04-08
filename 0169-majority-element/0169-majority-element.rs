use std::collections::HashMap;

impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut counter = HashMap::<i32,i32>::new();
        let mut threshold = (nums.len() / 2) as i32;
        for num in nums.iter() {
            let count = counter.entry(*num).or_insert(0);
            *count += 1;
        }
        let mut result: i32 = 0;
        
        for (k, v) in counter {
            if v > threshold {
                result = k;
                break;
            }
        }

        return result;
    }
}