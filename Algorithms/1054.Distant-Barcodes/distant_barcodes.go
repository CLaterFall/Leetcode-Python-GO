type Pair struct {
    Key   int
    Value int
}
type PairList []Pair

func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func rearrangeBarcodes(barcodes []int) []int {
    cnt_map := make(map[int]int)
    for i := 0; i < len(barcodes); i++ {
        cnt_map[barcodes[i]]++
    }
    
    pair_list := make([]Pair, 0, len(cnt_map))
	for k, v := range cnt_map {
		pair_list = append(pair_list, Pair{k, v})
	}
	sort.Slice(pair_list, func(i, j int) bool {
		return pair_list[i].Value > pair_list[j].Value
	})

    res := make([]int, len(barcodes))
    start := 0
	for _, v := range pair_list {
		val, count := v.Key, v.Value
		for count > 0 {
			res[start] = val
			start += 2
			count--
			if start >= len(barcodes) {
				start = 1
			}
		}
	}
    
    return res
}