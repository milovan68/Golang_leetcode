func removeElement(nums []int, val int) int {
    var idx int
    for idx < len(nums){
        if nums[idx] == val {
            nums = append(nums[:idx], nums[idx+1:]...)
        } else {
            idx ++ 
        }
    }
    return len(nums)
}