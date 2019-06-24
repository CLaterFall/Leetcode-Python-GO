func dailyTemperatures(T []int) []int {
    length := len(T)
    res := make([]int, length)
    stack := make([]int, 0)
    
    for i := 0; i < length; i++ {
        slen := len(stack)
        for slen > 0 && T[i] > T[stack[slen - 1]] {
            res[stack[slen - 1]] = i - stack[slen - 1]
            stack = stack[:slen - 1]
            slen -= 1
        }
        stack = append(stack, i)
    }
    
    return res        
}