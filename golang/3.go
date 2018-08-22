func lengthOfLongestSubstring(s string) int {
    var max int = 0
    var l int = 0
    var b int = 0
    if s == "" {
        return 0
    }
    for i := 1; i < len(s); i++ {
        e := strings.IndexByte(s[b:i], s[i])
        if e < 0 {
            l += 1
        } else {
            b = e + b + 1
            l = i - b
        }
        if l > max {
            max = l
        }
    }
    return max + 1
}