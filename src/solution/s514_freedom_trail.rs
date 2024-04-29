struct Solution{}

impl Solution {
    /*
        General idea:
        - 
     */
    pub fn find_rotate_steps(ring: String, key: String) -> i32 {
        let mut track: [Vec<i32>;26] = [Vec::new();26];
    }
}

#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn test_cases() {
        assert_eq!(4, Solution::find_rotate_steps("godding".to_string(), "gd".to_string()));
        assert_eq!(13, Solution::find_rotate_steps("godding".to_string(), "godding".to_string()));
        
    }
}
