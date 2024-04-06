impl Solution {
    pub fn is_perfect_square(num: i32) -> bool {
        let mut left: i64 = 1;
        let mut right: i64 = num as i64;

        let mut mid: i64 =( (right - left)/2 + left) as i64;
        let mut tmp: i64 = 0;
        // binary search
        // while left <= right
        let num64 = num as i64;
        while left >= 0 &&  left <= right {
            mid = ( (right - left)/2 + left) as i64;
            tmp = mid * mid;
            if tmp == num64 {
                return true
            }else if tmp >= 0 && tmp < num64 {
                left = mid + 1;
            }else {
                right = mid - 1;
            }
            // println!("left={}, right={}, mid={}, tmp={}", left, right, mid, tmp)

        }
        return false
    }
}