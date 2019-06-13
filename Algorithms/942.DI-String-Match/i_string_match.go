func diStringMatch(S string) []int {
    res := []int{}
    start := 0
    end := len(S)
    for _, ele := range S {
        if string(ele) == "I" {
            res = append(res, start)
            start += 1
        } else {
            res = append(res, end)
            end -= 1
        }
    }
    res = append(res, start)
    return res
}
