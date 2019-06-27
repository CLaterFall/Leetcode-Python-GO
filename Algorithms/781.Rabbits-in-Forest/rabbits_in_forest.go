func numRabbits(answers []int) int {
    count := [1000]int{}
    for _, ele := range answers {
        count[ele] += 1
    }
    res := 0
    for ans, cnt := range count {
        if cnt == 0 {
            continue
        }
        ans ++
        res += cnt / ans * ans
        if cnt % ans > 0 {
            res += ans
        }
    }
    
    return res
}