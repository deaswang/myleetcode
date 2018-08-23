func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    var l = len(s)
    var max = 1
    var ret = s[0:1]
    var j, k int
    for i := 0; i < l; i++ {
        j, k = i-1, i+1
        for j>=0 && k < l && s[j] == s[k] {
            j--
            k++
        }
        if k-j-1 > max {
            max = k-j-1
            ret = s[j+1:k]
        }
        j, k = i, i+1
        for j>=0 && k < l && s[j] == s[k] {
            j--
            k++
        }
        if k-j-1 > max {
            max = k-j-1
            ret = s[j+1:k]
        }
    }
    return ret
}
