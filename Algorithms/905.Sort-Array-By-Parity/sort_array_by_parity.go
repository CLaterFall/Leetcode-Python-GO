func sortArrayByParity(A []int) []int {
	var res []int
	var odd []int
	for _, ele := range A {
		if ele%2 == 0 {
			res = append(res, ele)
		} else {
			odd = append(odd, ele)
		}
	}
	res = append(res, odd...)
	return res
}
