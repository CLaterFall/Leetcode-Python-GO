func minDeletionSize(A []string) int {
	if len(A) == 0 {
		return 0
	}
	length := len(A[0])
	res := 0
	for j := 0; j < length; j++ {
		for i := 1; i < len(A); i++ {
			if A[i][j] < A[i-1][j] {
				res++
				break
			}
		}
	}
	return res
}
