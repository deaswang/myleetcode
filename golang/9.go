func isPalindrome(x int) bool {
    if x<0 {
        return false
    }
    var num = x
    var nums = make([]int, 0)
    for num > 0 {
        nums = append(nums, num%10)
        num = num/10
    }
    for j, k := 0, len(nums)-1; j < k ; {
        if nums[j] != nums[k] {
            return false
        }
        j++
        k--
    }
    return true
}
