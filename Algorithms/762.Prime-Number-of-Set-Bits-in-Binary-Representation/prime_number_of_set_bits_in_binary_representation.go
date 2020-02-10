func countPrimeSetBits(L int, R int) int {
	res := 0
	for i := L; i <= R; i++ {
		count := 0
		for j := i; j > 0; j /= 2 {
			lsb := j % 2
			if lsb == 1 {
				count += 1
			}
		}
		for _, ele := range []int{2, 3, 5, 7, 11, 13, 17, 19} {
			if ele == count {
				res += 1
			}
		}
	}
	return res
}
