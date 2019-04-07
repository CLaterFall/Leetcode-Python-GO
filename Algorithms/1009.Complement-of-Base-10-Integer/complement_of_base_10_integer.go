func bitwiseComplement(N int) int {
    if N == 0 {
        return 1
    }
    res := 0
    base := 1
    for {
        if N <= 0 {
            break
        }
        if N % 2 == 0 {
            res += base
        }
        base *= 2
        N /= 2
    }
    return res
}