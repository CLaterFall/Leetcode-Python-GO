func isMonotonic(A []int) bool {
	flag := 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			continue
		}

		cur_flag := 0
		if A[i] < A[i+1] {
			cur_flag = 1
		} else {
			cur_flag = 2
		}

		if flag == 0 {
			flag = cur_flag
			continue
		}
		if flag != cur_flag {
			return false
		}
	}
	return true
}
