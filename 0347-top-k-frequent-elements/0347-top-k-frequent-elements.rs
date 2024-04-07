use std::collections::BinaryHeap;
use std::collections::HashMap;
use std::cmp::Ordering;
use std::cmp::Reverse;
use std::clone::Clone;


impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut counter = HashMap::<i32,i32>::new();
        let mut result: Vec<i32> = vec![];

        for num in nums {
            if counter.contains_key(&num){
                counter.insert(num, counter.get(&num).unwrap()+1);
            }else {
                counter.insert(num, 1);
            }
        }
        let mut minHeap = BinaryHeap::new();
        for (key, value) in counter {
            if (minHeap.len() as i32) < k {
                minHeap.push(Reverse(RecordAscOrder{num: key, freq: value}));
            }else {
                let top = minHeap.peek().unwrap().0;
                if top.freq < value {
                    minHeap.pop();
                    minHeap.push(Reverse(RecordAscOrder{num: key, freq: value}));
                }
            }
        }
        for rec in minHeap {
            result.push(rec.0.num);
        }
        return result;
    }
}


#[derive(PartialEq, Eq, Copy)]
struct RecordAscOrder {
    num: i32,
    freq: i32,
}

impl Clone for RecordAscOrder {
    fn clone(&self) -> Self {
        return *self;
    }
}

impl Ord for RecordAscOrder {
    fn cmp(&self, other: &Self) -> Ordering {
        return self.freq.cmp(&other.freq)
    }
}

impl PartialOrd for RecordAscOrder {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        return Some(self.cmp(&other))
    }
}