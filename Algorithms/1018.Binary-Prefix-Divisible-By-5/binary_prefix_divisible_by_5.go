import "fmt"
func prefixesDivBy5(A []int) []bool {
    length := len(A)
    i := 0
    res := make([]bool, length)
    for i = 1; i < length; i++ {
        A[i] += A[i - 1] * 2 % 5
    }
    
    for i = 0; i < length; i++ {
        if (A[i] % 5 == 0) {
            res[i] = true
        }
    }
    return res
}
