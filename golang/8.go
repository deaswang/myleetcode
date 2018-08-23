func myAtoi(str string) int {
    var ret int64
    var sign = true
    s := strings.TrimLeft(str, " ")
    if s == "" {
        return 0
    }
    var i = 0
    if s[i] >= '0' && s[i] <= '9' {
        i = 0
    } else if s[i] == '-' && i+1<len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
        sign = false
        i++
    } else if s[i] == '+' && i+1<len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
        sign = true
        i++
    } else {
        return 0
    }
    for ; i<len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            ret = 10 * ret + int64(s[i] - '0')
            if ret > 2147483648 {
                break
            }
        } else {
            break
        }
    }
    if !sign {
        ret = -ret
    }
    if ret > 2147483647 {
        return 2147483647
    }
    if ret < -2147483648 {
        return -2147483648
    }
    return int(ret)
}
