impl Solution {
    pub fn check_valid_string(s: String) -> bool {
        return !Solution::isOverOpenBracket(&s) && !Solution::isOverClosedBracket(&s)
    }

    fn isOverClosedBracket(s: &String) -> bool {
        let mut balance = 0;
        let mut wild = 0;

        for b in s.chars() {
            if b == '(' {
                balance += 1;
            }else if b == ')' {
                if balance > 0 {
                    balance -= 1;
                }else if wild > 0 {
                    wild -= 1
                }else{
                    return true
                }
            }else if b == '*' {
                wild += 1;
            }
        }
        return false
    }

    fn isOverOpenBracket(s: &String) -> bool {
        let mut balance = 0;
        let mut wild = 0;

        for b in s.chars().rev() {
            if b == ')' {
                balance += 1;
            }else if b == '(' {
                if balance > 0 {
                    balance -= 1;
                }else if wild > 0 {
                    wild -= 1
                }else{
                    return true
                }
            }else if b == '*' {
                wild += 1;
            }
        }
        return false
    }
}