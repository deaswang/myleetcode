func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, n := range nums {
        m[n] = i
    }
    for i, n := range nums {
        if j, ok := m[target - n]; ok {
            if j != i {
                return []int{i, j}
            }
        }
    }
    return []int{}
}