struct Solution {}

impl Solution {
  const  UPPER: usize = 37;
  const TRIBONACCI_NUMS: [i32;Self::UPPER+1] = Solution::tribonacci_const();

  pub fn tribonacci(n: i32) -> i32 {
    return Solution::TRIBONACCI_NUMS[n as usize];
  }

  const fn tribonacci_const() -> [i32;Self::UPPER+1]{
    let mut tribo = [0;Self::UPPER+1];
    tribo[0] = 0;
    tribo[1] = 1;
    tribo[2] = 1;

    let mut i = 3;
    loop {
        if i == Self::UPPER+1 {
          break;
        }
        tribo[i] = tribo[i-1] + tribo[i-2] + tribo[i-3];
        i += 1;
    }
    return tribo;
  }
}  

#[cfg(test)]
mod test {
    use super::Solution;

  #[test]
  fn run_test() {
    assert_eq!(4, Solution::tribonacci(4));
    assert_eq!(1389537, Solution::tribonacci(25));
  }
}
