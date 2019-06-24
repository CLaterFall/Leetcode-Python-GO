func customSortString(S string, T string) string {
    res := ""
    length := len(S)
    for i := 0; i < length; i++ {
        cnt := strings.Count(T, S[i:i+1])
        res += strings.Repeat(S[i:i+1], cnt)
        T = strings.Replace(T, S[i:i+1], "", -1)
    }
    return res + T
}