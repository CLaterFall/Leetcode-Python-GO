func peakIndexInMountainArray(A []int) int {
    for idx, value := range A {
        if value > A[idx + 1] {
            return idx
        }
    }
    return 0
}