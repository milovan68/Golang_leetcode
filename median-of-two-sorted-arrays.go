func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var countNums1, countNums2 int
    len1 := len(nums1)
    len2 := len(nums2)
    resArray := make([]int, (len1+len2) / 2 + 1)
    len3 := len(resArray)
    for i := 0; i < len3; i ++ {
        if countNums1 == len(nums1) {
            resArray[i] = nums2[countNums2]
            countNums2 ++
        } else if countNums2 == len(nums2){
            resArray[i] = nums1[countNums1]
            countNums1 ++
        } else {
            if nums1[countNums1] >= nums2[countNums2]{
                resArray[i] = nums2[countNums2]
                countNums2 ++
            } else {
                resArray[i] = nums1[countNums1]
                countNums1 ++
            }
        }
    }
    if (len1 + len2) % 2 == 0 {
        return (float64(resArray[len3 - 1]) + float64(resArray[len3 - 2])) / 2
    } else {
        return float64(resArray[len3 - 1])
    }
}