import "strings"

func reformat(s string) string {
	var str_arr []string
	var digit_arr []string
	for _, ele := range s {
		if 48 <= ele && ele <= 58 {
			digit_arr = append(digit_arr, string(ele))
		} else {
			str_arr = append(str_arr, string(ele))
		}
	}
	lend := len(digit_arr)
	lens := len(str_arr)
	len_diff := lend - lens
	if len_diff > 1 || len_diff < -1 {
		return ""
	}
	res := make([]string, lend+lens)
	pos := 0
	if lend > lens {
		pos = 0
	} else {
		pos = 1
	}
	for _, ele := range digit_arr {
		res[pos] = ele
		pos += 2
	}
	if lend > lens {
		pos = 1
	} else {
		pos = 0
	}
	for _, ele := range str_arr {
		res[pos] = ele
		pos += 2
	}
	return strings.Join(res, "")
}
