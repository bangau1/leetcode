impl Solution {
    pub fn count_students(mut students: Vec<i32>, mut sandwiches: Vec<i32>) -> i32 {
        let mut pref: [usize; 2] = [0, 0];

        for st in students.iter() {
            pref[*st as usize] += 1;
        }

        while students.len() > 0 && pref[sandwiches[0] as usize] > 0 {
            
            while students.len() > 0 && students[0] == sandwiches[0] {
                pref[students[0] as usize] -= 1;
                students.remove(0);
                sandwiches.remove(0);
            }

            while sandwiches.len() > 0 && students[0] != sandwiches[0] && pref[sandwiches[0] as usize] > 0 {
                let front = students.remove(0);
                students.push(front);
            }
        }

        return students.len() as i32;
    }
}