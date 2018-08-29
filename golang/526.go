func countArrangement(N int) int {
    var arr = make([]int, N)
    return placeN(arr, N)
}

func placeN(arr []int, n int) int {
    if n <= 1 {
        return 1
    }
    var ret = 0
    for i:=0; i<len(arr) ; i++ {
        if arr[i] == 0 {
            if (i+1)%n == 0 || n%(i+1) == 0 {
                newarr := make([]int, len(arr))
                copy(newarr, arr)
                newarr[i] = n
                ret += placeN(newarr, n-1)
            } 
        }
    }
    return ret
}

