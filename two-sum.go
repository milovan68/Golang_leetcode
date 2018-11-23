func twoSum(nums []int, target int) []int {
    diff := make(map[int]int)
    for idx, value := range nums{
        if pos, ok := diff[target - value]; ok{
            if pos != idx {
                return []int{pos, idx}
            }
        } else {
            diff[value] = idx
        }
    }
    return nil
}