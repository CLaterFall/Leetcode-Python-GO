import "strings"

func commonChars(A []string) []string {
	t := []rune(A[0])
	temp := make(map[string]int)
	for i := 0; i < len(t); i++ {
		temp[string(t[i])] = strings.Count(A[0], string(t[i]))
	}

	for i := 1; i < len(A); i++ {
		for key, value := range temp {
			cnt := strings.Count(A[i], string(key))
			if cnt < value {
				temp[string(key)] = cnt
			}
		}
	}
	var res []string
	for key, value := range temp {
		for i := 0; i < value; i++ {
			res = append(res, key)
		}
	}
	return res
}
