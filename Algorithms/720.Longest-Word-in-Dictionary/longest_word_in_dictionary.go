import "sort"

func longestWord(words []string) string {
	sort.Strings(words)
	res := words[0]
	length := len(words)
	var temp = make(map[string]bool, length)
	for _, word := range words {
		word_len := len(word)
		if word_len == 1 {
			temp[word] = true
		} else if temp[word[:word_len-1]] {
			temp[word] = true
			if len(res) < word_len {
				res = word
			}
		}
	}
	return res
}
