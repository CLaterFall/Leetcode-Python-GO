import "sort"

func sortedSquares(A []int) []int {
	var res []int
	for _, ele := range A {
		res = append(res, ele*ele)
	}
	sort.Ints(res)
	return res
}
