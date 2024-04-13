impl Solution {
    pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
        let mut max_area: i32 = 0;
        let mut stack: Vec<usize> = vec![];

        for i in 0..heights.len() {
            while stack.len() > 0 && heights[i] < heights[*stack.last().unwrap()] {
                let curr_h = heights[stack.pop().unwrap()];
                let mut curr_w = i as i32;
                if stack.len() > 0 {
                    curr_w = curr_w - (*stack.last().unwrap() as i32 + 1);
                }
                let area = curr_h * curr_w;
                if max_area < area {
                    max_area = area;
                }
            }
            stack.push(i);
        }
        while stack.len() > 0 {
            let curr_h = heights[stack.pop().unwrap()];
            let mut curr_w = heights.len() as i32;
            if stack.len() > 0 {
                curr_w = curr_w - (*stack.last().unwrap() as i32 + 1);
            }
            let area = curr_h * curr_w;
            if max_area < area {
                max_area = area;
            }
        }
        return max_area;   
    }
}