import "sort"

func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	total := 0
	for _, ele := range nums {
		total = total + ele
	}
	var res []int
	cur := 0
	for _, ele := range nums {
		cur = cur + ele
		if cur > total-cur {
			res = append(res, ele)
			return res
		}
		res = append(res, ele)
	}
	return res
}
