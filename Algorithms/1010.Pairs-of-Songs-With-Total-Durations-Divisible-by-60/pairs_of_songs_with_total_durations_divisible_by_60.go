func numPairsDivisibleBy60(time []int) int {
    res := 0
    cntMap := make([]int, 60)
    for _, ele := range time {
        res += cntMap[(60 - ele % 60) % 60]
        cntMap[ele % 60] += 1
    }
    return res
}