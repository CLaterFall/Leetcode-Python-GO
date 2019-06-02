func selfDividingNumbers(left int, right int) []int {
    res := []int{}
    for i := left; i <= right; i++ {
        temp := i
        flag := true
        for temp > 0 {
            remain := temp % 10
            temp = temp / 10
            if remain == 0 || i % remain > 0 {
                flag = false
                break
            }   
        }
        if flag {
            res = append(res, i)
        }
    }
    return res
}