func addToArrayForm(A []int, K int) []int {
	var temp int
	var res []int
	length := len(A) - 1
	for i := length; i >= 0; i-- {
		ele := A[i]
		if K > 0 {
			temp = K%10 + ele
			if temp >= 10 {
				res = append(res, temp%10)
				K = K/10 + temp/10
			} else {
				res = append(res, temp)
				K = K / 10
			}
		} else {
			res = append(res, ele)
		}
	}
	for i := K; i > 0; i /= 10 {
		res = append(res, i%10)
	}

	var ret []int
	for i := len(res) - 1; i >= 0; i-- {
		ret = append(ret, res[i])
	}
	return ret
}
