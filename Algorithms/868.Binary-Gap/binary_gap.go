func binaryGap(N int) int {
    for N != 0 && N & 1 != 1 {
        N = N >> 1
    }
    
    max_dis := 0
    consecutive := 0
    for N > 0 {
        N = N / 2
        consecutive += 1
        if N & 1 == 1 {
            if max_dis < consecutive {
                max_dis = consecutive
            }
            consecutive = 0
        }
    }
    return max_dis
}