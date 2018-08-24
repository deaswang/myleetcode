func isSelfCrossing(x []int) bool {
    if len(x) < 4 {
        return false
    }
    if len(x) == 4 {
        if x[3] >= x[1] && x[2] <= x[0] {
            return true
        }
    }
    if len(x) == 5 {
        if (x[4] >= x[2] && x[3] <= x[1]) || (x[4] + x[0] >= x[2] && x[3] == x[1]) {
            return true
        }
    }
    for i:= 5; i<len(x) ; i++ {
        if x[i] >= x[i-2] && x[i-1] <= x[i-3]  || (x[i] + x[i-4] >= x[i-2] && x[i-1] == x[i-3]) || (x[i-2] >= x[i-4] && x[i-3] >= x[i-1] && x[i-1] >= x[i-3] - x[i-5] && x[i] >= x[i-2] - x[i-4]) {
            return true
        }
    }
    return false
}
