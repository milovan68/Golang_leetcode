func removeDuplicates(nums []int) int {
    idx := 1
    for idx < len(nums) {
        if nums[idx] == nums[idx - 1] {
            nums = append(nums[:idx], nums[idx+1:]...)
        } else {
            idx ++
        }
    }
    return len(nums)
}