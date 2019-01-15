func findLength(A []int, B []int) int {
	lenA, lenB := len(A), len(B)
	dp := make([][]int, lenA)
	for i := range dp {
		dp[i] = make([]int, lenB)
	}

	res := 0
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			if A[i] == B[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
