func numJewelsInStones(J string, S string) int {
    var ret = 0
    var js = make(map[rune]bool)
    for _, j := range J {
        js[j] = true
    }
    for _, s := range S {
        if _, ok := js[s]; ok {
            ret++
        }
    }
    return ret
}
