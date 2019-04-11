import "math"
func powerfulIntegers(x int, y int, bound int) []int {
    var res []int
    if bound <= 1 {
        return res
    }
    lgx := 0
    if x > 2 {
        lgx = int(math.Log(float64(bound)) / math.Log(float64(x))) + 1
    } else {
        lgx = int(math.Log(float64(bound)) / math.Log(float64(2))) + 1
    }
    lgy := 0
    if y > 2 {
        lgy = int(math.Log(float64(bound)) / math.Log(float64(y))) + 1
    } else {
        lgy = int(math.Log(float64(bound)) / math.Log(float64(2))) + 1
    }
    
    i := 0
    resMap := make(map[int]bool)
    fmt.Println(lgx,lgy)
    for i < lgx {
        j := 0
        for j < lgy {
            ssum := int(math.Pow(float64(x), float64(i)) + math.Pow(float64(y), float64(j)))
            if ssum > bound {
                break
            }
            resMap[ssum] = true
            j++
        }
        i++
    }
    for k,_ := range resMap {
        res = append(res, k)
    }
    return res
}