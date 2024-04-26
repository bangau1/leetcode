
use std::cmp::{max, min};

impl Solution {
    pub fn longest_ideal_string(s: String, k: i32) -> i32 {
        let mut track: [i32; 26] = [0; 26];
        let mut longest: i32 = 0;
        let s = s.as_bytes();

        for (i, &c) in s.iter().enumerate() {
            let current_char = (c - 'a' as u8) as i32;
            let mut prev_max = 0;

            let (left, right) = (
                max(0, current_char - k) as usize,
                min(25, current_char + k) as usize,
            );

            for prev in left..right + 1 {
                if prev_max < track[prev] {
                    prev_max = track[prev];
                }
                if prev_max == i as i32 {
                    break;
                }
            }
            track[current_char as usize] = prev_max + 1;
            if longest < track[current_char as usize] {
                longest = track[current_char as usize];
            }
        }

        return longest;
    }
}
