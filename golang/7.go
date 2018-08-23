func reverse(x int) int {
    var ret int64
    var sign = true
    var num int64 = int64(x)
    var i int
    if num < 0 {
        sign = false
        num = -num
    }
    var nums = make([]int64, 10)
    for i = 0 ; num > 0; i++ {
        nums[i] = num%10
        num = num/10
    }
    for j := 0 ; j < i; j++ {
        ret = 10*ret + nums[j]
    }
    if !sign {
        ret = -ret
    }
    if ret > 2147483647 || ret < -2147483648 {
        return 0
    }
    return int(ret)
}
