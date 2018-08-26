func longestSubstring(s string, k int) int {
    var count = make(map[rune]int)
    for _, c := range s {
        if _, ok := count[c]; ok {
            count[c]++
        } else {
            count[c] = 1
        }
    }
    for b, v:= range count {
        if v < k {
            ss := strings.Split(s, string(b))
            ret := 0
            for _, sss := range ss {
                value := longestSubstring(sss, k)
                if ret <= value {
                    ret = value
                }
            }
            return ret
        }
    }
    return len(s)
}