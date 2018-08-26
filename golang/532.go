func findPairs(nums []int, k int) int {
    var pair = make(map[int]int)
    var a, b int
    for i, n := range nums {
        for _, m := range nums[i+1:] {
            if n < m {
                a, b = m, n
            } else {
                b, a = m, n
            }
            if a-b == k {
                if _, ok := pair[a]; !ok {
                    pair[a] = b
                }
            }
        }
    }
    return len(pair)
}

// better
func findPairs(nums []int, k int) int {
    if k<0 {
        return 0
    }
    var pair = make(map[int]int)
    if k == 0{
        for _, n := range nums {
            if _, ok := pair[n]; ok {
                pair[n]++
            } else {
                pair[n] = 0
            }
        }
        ret := 0
        for _, pv := range pair {
            if pv > 0 {
                ret++
            }
        }
        return ret
    }
    for _, n := range nums {
        if _, ok := pair[n]; !ok {
            pair[n] = n-k
        }
    }
    ret := 0
    for _, pv := range pair {
        if _, ok := pair[pv]; ok{
            ret++
        }
    }
    return ret
}
