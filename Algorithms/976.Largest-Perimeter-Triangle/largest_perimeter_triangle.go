import "sort"

func largestPerimeter(A []int) int {
	length := len(A)
	sort.Ints(A)
	for i := length - 1; i >= 2; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}
