func buddyStrings(A string, B string) bool {
	lenA, lenB := len(A), len(B)
	if lenA == 0 || lenB == 0 || lenA != lenB {
		return false
	}

	if A == B {
		s := make(map[rune]int, lenA)
		for i, ele := range A {
			s[ele] = i
		}
		return len(s) < lenA
	}

	i := 0
	cnt := 0
	diff := make([]int, 2)
	total := 0
	for ; i < lenA; i++ {
		if A[i] != B[i] {
			if cnt < 2 {
				diff[cnt] = i
				cnt++
			}
			total++
		}
	}

	return (total == 2 && A[diff[1]] == B[diff[0]] && A[diff[0]] == B[diff[1]])
}
