func lemonadeChange(bills []int) bool {
    five_cnt, ten_cnt := 0, 0
    for _, value := range bills {
        switch value {
            case 5:
            five_cnt++
            case 10:
            five_cnt--
            ten_cnt++
            case 20:
            if ten_cnt > 0 {
                ten_cnt--
                five_cnt--
            } else {
                five_cnt -= 3
            }
        }
        if five_cnt < 0 || ten_cnt < 0 {
            return false
        }
    }
    return true
}
