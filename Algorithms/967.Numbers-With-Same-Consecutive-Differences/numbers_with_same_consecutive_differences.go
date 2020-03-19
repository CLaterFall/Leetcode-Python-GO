func numsSameConsecDiff(N int, K int) []int {
    num := []int{0, 1,2,3,4,5,6,7,8,9}
    for i:=0; i< N-1; i++ {
        tmp := []int{}
        for _, v := range num {
            if v == 0 && i == 0{
                continue
            }
            reminder := v % 10
            if reminder + K < 10 {
                tmp = append(tmp, v*10 + reminder + K)
            }
            if K > 0 && reminder - K >= 0 {
                tmp = append(tmp, v*10 + reminder - K)
            }
        }
        num = tmp
    }
    return num
}
