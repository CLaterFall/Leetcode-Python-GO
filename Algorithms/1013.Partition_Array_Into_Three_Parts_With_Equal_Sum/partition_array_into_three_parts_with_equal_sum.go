func canThreePartsEqualSum(A []int) bool {
    total := 0
    for _, value := range A {
        total += value
    }
    
    if total % 3 != 0 {
        return false
    }
    count := 0
    part := total / 3
    cur := 0
    for _, value := range A {
        cur += value
        if cur == part {
            count ++
            cur = 0
        }
    }
    if count == 3 {
        return true
    }
    
    
    return false
}