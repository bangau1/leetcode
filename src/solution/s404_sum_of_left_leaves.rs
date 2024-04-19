pub struct Solution {}
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
use std::cell::RefCell;
use std::rc::Rc;

impl Solution {
    pub fn sum_of_left_leaves(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        // if root is not None
        if let Some(r) = root {
            let r = r.borrow();
            let mut sum = 0;

            // if root.Left is not none
            if let Some(rl) = &r.left {
                let rl = rl.borrow();
                // if root.left is the leave
                if rl.left.is_none() && rl.right.is_none() {
                    sum += rl.val;
                } else {
                    sum += Solution::sum_of_left_leaves(r.left.clone());
                }
            }

            if r.right.is_some() {
                sum += Solution::sum_of_left_leaves(r.right.clone())
            }

            return sum;
        } else {
            return 0;
        }
    }
}
