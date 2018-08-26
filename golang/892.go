func surfaceArea(grid [][]int) int {
    var N = len(grid)
    var ret = N*N*2
    for i := 0; i<N ; i++ {
        ret+=grid[i][0]+grid[i][N-1]+grid[0][i]+grid[N-1][i]
    }
    for i := 1; i<N; i++ {
        for j:=0;j<N; j++ {
            if grid[i][j] > grid[i-1][j] {
                ret+=grid[i][j] - grid[i-1][j]
            } else {
                ret+=grid[i-1][j] - grid[i][j]
            }
        }
    }
    for i := 0; i<N; i++ {
        for j:=1;j<N; j++ {
            if grid[i][j] > grid[i][j-1] {
                ret+=grid[i][j] - grid[i][j-1]
            } else {
                ret+=grid[i][j-1] - grid[i][j]
            }
        }
    }
    var zero = 0
    for i := 0; i<N; i++ {
        for j:=0;j<N; j++ {
            if grid[i][j] == 0 {
                zero++
            }
        }
    }
    ret-=2*zero
    return ret
}
