func minimumDeleteSum(s1 string, s2 string) int {
	len1, len2 := len(s1), len(s2)
	dp := make([]int, len2+1)
	for i := range dp {
		dp[i] = 0
	}

	for i := 1; i < len2+1; i++ {
		dp[i] = dp[i-1] + int(s2[i-1])
	}

	for i := 1; i < len1+1; i++ {
		temp_j_one := dp[0]
		dp[0] += int(s1[i-1])
		for j := 1; j < len2+1; j++ {
			temp_j := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = temp_j_one
			} else {
				dp[j] = min(temp_j+int(s1[i-1]), dp[j-1]+int(s2[j-1]))
			}
			temp_j_one = temp_j
		}
	}
	return dp[len2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
