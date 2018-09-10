func thirdMax(nums []int) int {
    var max = make([]int, 3)
    var empty = 3
    for _, i := range nums {
        j := 2 - empty
        for ; j>-1; j-- {
            if max[j] >= i {
                break
            }
        }
        if j > -1 && max[j] == i || j>1 {
            continue
        }
        j++
        temp := max[j]
        max[j] = i
        if empty>0 {
            empty--
        }
        for j++; j<3-empty ; j++ {
            max[j], temp = temp, max[j]
        }
    }
    if empty > 0 {
        return max[0]
    } else {
        return max[2]
    }
}
