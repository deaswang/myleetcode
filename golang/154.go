func findMin(nums []int) int {
    var l = len(nums)
    var i = 0
    if nums[i] >= nums[l-1] {
        i = i + (l-i)/2
        for nums[i] >= nums[l-1] && i < l-1 {
            i = i + (l-i)/2
        }
        for ; i>0; i-- {
            if nums[i] < nums[i-1] {
                return nums[i]
            }
        }
        return nums[i]
    } else {
        return nums[0]
    }
    return 0
}
