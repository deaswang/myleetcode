 
func countSmaller(nums []int) []int {
    var ret = make([]int, len(nums))
    for i := 0; i< len(nums); i++ {
        for j:= i+1; j< len(nums); j++ {
            if nums[i]>nums[j] {
                ret[i]++
            }
        }
    }
    return ret
}
