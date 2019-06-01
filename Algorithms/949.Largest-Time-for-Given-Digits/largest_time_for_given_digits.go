var permutations = [][]int{
	{0, 1, 2, 3},
	{0, 1, 3, 2},
	{0, 2, 1, 3},
	{0, 2, 3, 1},
	{0, 3, 1, 2},
	{0, 3, 2, 1},
	{1, 0, 2, 3},
	{1, 0, 3, 2},
	{1, 2, 0, 3},
	{1, 2, 3, 0},
	{1, 3, 0, 2},
	{1, 3, 2, 0},
	{2, 0, 1, 3},
	{2, 0, 3, 1},
	{2, 1, 0, 3},
	{2, 1, 3, 0},
	{2, 3, 1, 0},
	{2, 3, 0, 1},
	{3, 0, 2, 1},
	{3, 0, 1, 2},
	{3, 1, 0, 2},
	{3, 1, 2, 0},
	{3, 2, 0, 1},
	{3, 2, 1, 0},
}
func largestTimeFromDigits(A []int) string {
    hour, min := -1, -1
    res := ""
    for _, p := range permutations {
        p1, p2, p3, p4 := A[p[0]], A[p[1]], A[p[2]], A[p[3]]
        temp_h, temp_m := p1*10 + p2, p3*10 + p4
        if temp_h > 23 || temp_m > 59 || temp_h*60+temp_m < hour*60+min {
            continue
        }
        hour, min = temp_h, temp_m
    }
    if hour != -1 && min != -2 {
        res = fmt.Sprintf("%02d:%02d", hour, min)
    }
    
    return res
}