impl Solution {
    pub fn time_required_to_buy(tickets: Vec<i32>, k: i32) -> i32 {
        let mut people: Vec<[i32;2]> = Vec::new();

        for (idx , &el) in tickets.iter().enumerate() {
            if idx == k as usize {
                people.push([el, 1]);
            }else {
                people.push([el, 0]);   
            }
        }
        let mut temp: [i32;2];
        let mut time = 0;
        while people.len() > 0 {
            time += 1;
            if people[0][0] - 1 == 0 {
                if people[0][1] == 1 {
                    break
                }else{
                    people.remove(0);
                }
            }else {
                temp = people[0];
                temp[0] -= 1;

                people.remove(0);
                people.push(temp);
            }
        }

        return time
    }
}