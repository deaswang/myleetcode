func bulbSwitch(n int) int {
    var ret int
    var k, j int
    for i := 1; i<n+1; i++ {
        k = 0
        for j = 1; j <= int(math.Sqrt(float64(i))); j++ {
            if i%j == 0 {
                k+=2
            }
            if j*j == i {
                k-=1
            }
        }
        if k%2 == 1 {
            ret++
        }
    }
    return ret
}

// up is Time Limit Exceeded but only when j*j == i, then k%2 == 1, so all i*i bulb is on
func bulbSwitch(n int) int {
    return int(math.Sqrt(float64(n)))
}
