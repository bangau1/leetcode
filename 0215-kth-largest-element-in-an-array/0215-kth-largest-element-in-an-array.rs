use std::{cmp::Reverse, collections::BinaryHeap};


impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let k = k as usize;
        let mut min_heap = BinaryHeap::new();
        for num in nums {
            if min_heap.len() < k {
                min_heap.push(Reverse(num));
            } else if let Some(&Reverse(min)) = min_heap.peek() {
                if min < num {
                    min_heap.pop();
                    min_heap.push(Reverse(num));
                }
            }
        }
        let &Reverse(kth) = min_heap.peek().unwrap();
        return kth;
    }
}