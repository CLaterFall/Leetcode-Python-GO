func countNegatives(grid [][]int) int {
    n := len(grid[0])
    res := 0
    for _, ele := range grid {
        for j := 0; j < n; j++ {
            if ele[j] < 0 {
                res = res + n - j
                j = n
            }
        }
    }
    return res
}
