func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    var round = 2 * numRows - 2
    var l = len(s)
    var rl = 0
    var i, j = 0, 0
    var ret = make([]byte, l)
    for j = 0; j<l/round; j++ {
        ret[rl] = s[round*j]
        rl++
    }
    if l%round > 0 {
        ret[rl] = s[round*j]
        rl++
    }
    for i = 1; i < numRows - 1; i++ {
        for j = 0; j<l/round; j++ {
            ret[rl] = s[round*j+i]
            rl++
            ret[rl] = s[round*j+round-i]
            rl++
        }
        if l%round > i {
            ret[rl] = s[round*j+i]
            rl++
        }
        if l%round > round-i{
            ret[rl] = s[round*j+round-i]
            rl++
        }
    }
    for j = 0; j<l/round; j++ {
        ret[rl] = s[round*j + i]
        rl++
    }
    if l%round > i {
        ret[rl] = s[round*j + i]
        rl++
    }
    return string(ret)
}
