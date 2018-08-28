func minEatingSpeed(piles []int, H int) int {
    var total = 0
    for _, p := range piles {
        total += p
    }
    var ret = total/H
    if ret<=0 {
        ret = 1
    }
    var hour = -1
    for ; hour < 0 ; {
        hour = H
        for _, p := range piles {
            if p%ret == 0 {
                hour -= p/ret
            } else {
                hour -= p/ret+1
            }
            if hour < 0 {
                break
            }
        }
        ret++
    }
    return ret-1
}
