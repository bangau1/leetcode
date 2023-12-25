func twoSum(numbers []int, target int) []int {
   // use two pointer approach
   l, r := 0, len(numbers)-1

   for l < r {
       sum := numbers[l] + numbers[r]
       if sum < target {
           l++
       }else if sum > target {
           r--
       }else{
            break        
       }
   }
   return []int{l+1, r+1}
}