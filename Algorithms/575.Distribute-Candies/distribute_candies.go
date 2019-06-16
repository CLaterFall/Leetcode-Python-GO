func distributeCandies(candies []int) int {
    pairs := len(candies) / 2
    m := make(map[int]struct{}, 0)
    for _, v := range candies {
        m[v] = struct{}{}
        if len(m) == pairs {
            break
        }
    }
    return len(m)
}