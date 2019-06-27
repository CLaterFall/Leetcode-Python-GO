func numJewelsInStones(J string, S string) int {
    res := 0
    for _, ele := range J {
        res += strings.Count(S, string(ele))
    }
    return res
}