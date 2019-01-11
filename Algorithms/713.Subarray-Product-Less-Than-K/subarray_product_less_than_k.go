func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	length := len(nums)
	left, right := 0, 0
	prod := 1
	for right = 0; right < length; right++ {
		prod *= nums[right]
		for {
			if left <= right && prod >= k {
				prod = prod / nums[left]
				left += 1
			} else {
				break
			}
		}
		res += right - left + 1
	}
	return res
}
