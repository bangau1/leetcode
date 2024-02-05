func triangleType(nums []int) string {
    sort.Ints(nums)

    // using triangle inequality formula, given a, b, c, if all combination of sum of 2 sides > the other side, then it's triangle
    if nums[0] + nums[1] > nums[2] && nums[0] + nums[2] > nums[1] && nums[1] + nums[2] > nums[0] {
        if nums[0] == nums[1] && nums[1] == nums[2] {
            return "equilateral"
        }
        if nums[0] == nums[1] && nums[0]!=nums[2] || nums[0] != nums[1] && nums[1] == nums[2] {
            return "isosceles"
        }

        return "scalene"
    }else{
        return "none"
    }
    
}