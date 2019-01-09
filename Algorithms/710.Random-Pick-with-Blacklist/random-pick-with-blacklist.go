import (
	"math/rand"
	"sort"
)

type Solution struct {
	blackDic map[int]int
	length   int
}

func Constructor(N int, blacklist []int) Solution {
	blackDic := make(map[int]int, len(blacklist))
	length := N - 1

	sort.Sort(sort.Reverse(sort.IntSlice(blacklist)))

	for _, ele := range blacklist {
		t, ok := blackDic[length]
		if ok {
			blackDic[ele] = t
		} else {
			blackDic[ele] = length
		}
		length--
	}
	return Solution{
		length:   length,
		blackDic: blackDic,
	}

}

func (this *Solution) Pick() int {
	num := rand.Intn(this.length + 1)
	if t, ok := this.blackDic[num]; ok {
		return t
	}
	return num
}
