func pivotIndex(nums []int) int {
    left := 0
    total := 0
    for _, num := range nums {
        total += num
    }
    for index, num := range nums {
        if left == total - num - left {
            return index
        }
        left += num
    }
    return -1
}