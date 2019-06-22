func judgeCircle(moves string) bool {
    x := 0
    y := 0
    for _, ele := range moves{
        switch string(ele) {
            case "U":
            y += 1
            case "D":
            y -= 1
            case "L":
            x -= 1
            case "R":
            x += 1
        }
    }
    if x == 0 && y == 0 {
        return true
    } else {
        return false
    }
}