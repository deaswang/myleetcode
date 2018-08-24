func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    var length = len(matrix) * len(matrix[0])
    var ret = make([]int, length)
    var x, y = 0, 0
    var l, r, u, d = 0, len(matrix[0]) - 1, 0, len(matrix) - 1
    // direct [[0, 1], [1, 0], [0, -1], [-1, 0]]
    var dir = 0
    for i := 0; i < length; i++ {
        ret[i] = matrix[x][y]
        if dir == 0 {
            y++
            if y>=r {
                if y>r{
                    y--
                    x++
                }
                dir = 1
                u++
            }
        } else if dir == 1 {
            x++
            if x>=d {
                if x>d {
                    x--
                    y--
                }
                dir = 2
                r--
            }
        } else if dir == 2 {
            y--
            if y<=l {
                if y<l {
                    y++
                    x--
                }
                dir = 3
                d--
            }
        } else if dir == 3 {
            x--
            if x<=u {
                if x<u {
                    x++
                    y++
                }
                dir = 0
                l++
            }
        }
    }
    return ret
}
