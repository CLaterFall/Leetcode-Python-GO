import "strconv"

func calPoints(ops []string) int {
	res := 0
	pointStack := make([]int, 0, len(ops))
	for _, value := range ops {
		switch value {
		case "+":
			r1 := pointStack[len(pointStack)-1]
			r2 := pointStack[len(pointStack)-2]
			pointStack = append(pointStack, r1+r2)
			res = res + r1 + r2
		case "C":
			res = res - pointStack[len(pointStack)-1]
			pointStack = pointStack[:len(pointStack)-1]
		case "D":
			r1 := pointStack[len(pointStack)-1]
			pointStack = append(pointStack, 2*r1)
			res = res + r1*2
		default:
			num, _ := strconv.Atoi(value)
			pointStack = append(pointStack, num)
			res = res + num
		}
	}
	return res
}
