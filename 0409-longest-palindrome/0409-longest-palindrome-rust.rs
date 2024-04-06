use std::collections::HashMap;

impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        let mut counter = HashMap::<u8, i32>::new();
        let mut length = 0;
        let slen = s.len() as i32;
        
        for b in s.into_bytes() {
            match counter.get(&b) {
                Some(count) => {counter.insert(b, count + 1);},
                None => { counter.insert(b, 1);},
            }
        }

        for (k, v) in counter {
            if v % 2 == 0 {
                length += v
            }else{
                length += v-1
            }
        }

        if length < slen {
            length += 1;
        }
        return length        
    }
}