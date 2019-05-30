func largestSumAfterKNegations(A []int, K int) int {
    if K == 0 {
        return 0
    }
    min := 101
    sort.Ints(A)
    index := 0
    for K > 0 {
        if index >= len(A) {
            break
        }
        if A[index] >= 0 {
            if A[index] < min {
                min = A[index]
            }
            break
        }
        if A[index] < 0 {
            A[index] = -A[index]
        }
        if A[index] < min {
            min = A[index]
        }
        index++
        K--
    }
    
    total := 0
    for _, val := range A {
        total += val
    }
    if K % 2 == 0 {
        return total
    } else {
        return total - 2*min
    }
    
}