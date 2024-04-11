impl Solution {

    pub fn remove_kdigits(num: String, k: i32) -> String {
        let mut stack: Vec<u8> = vec![];
        let mut k = k;
        for b in num.as_bytes() {
            while stack.len() > 0 && k > 0 && stack.last().unwrap() > b {
                stack.pop();
                k-=1;
            }
            stack.push(*b);
        }
        
        while k > 0 && stack.len() > 0 {
            stack.pop();
            k-=1;
        }
        // remove trailing zero
        while stack.len() > 0 && *stack.first().unwrap() == "0".as_bytes()[0] {
            stack.remove(0);
        }
        if stack.len() == 0 {
            return String::from("0");
        }
        return String::from_utf8(stack).unwrap();
    }

}