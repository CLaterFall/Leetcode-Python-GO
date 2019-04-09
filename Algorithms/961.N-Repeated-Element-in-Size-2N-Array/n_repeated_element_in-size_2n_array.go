func repeatedNTimes(A []int) int {
    var ele [10000]bool
    for _, value := range A {
        if ele[value] {
            return value
        }
        ele[value] = true
    }
    return -1
}