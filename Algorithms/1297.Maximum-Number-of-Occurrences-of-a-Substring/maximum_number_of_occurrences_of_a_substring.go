func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
    length := len(s)
    if length < minSize {
        return 0
    }
    res := 0
    dic := make(map[string]int)
    for i := minSize; i<= length; i++ {
        tmp := s[i-minSize:i]
        
        flag := true
        cnt := make(map[rune]int)
        for _, ch := range tmp {
            _, ok := cnt[ch]
            if !ok {
                cnt[ch] = 1
            }
            if len(cnt) > maxLetters {
                flag = false
                break
            }
	    }

        if flag {
            _, ok := dic[tmp]
            if ok {
                dic[tmp] += 1    
            } else {
                dic[tmp] = 1
            }
            if dic[tmp] > res {
                res = dic[tmp]
            }
        }
    }
    return res
}
