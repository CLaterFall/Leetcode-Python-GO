func constructRectangle(area int) []int {
    W := int(math.Sqrt(float64(area)))
    for W >=1 && area % W != 0 {
        W -= 1
    }
    return []int{area/W, W}
}