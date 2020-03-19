func maxArea(height []int) int {
    head := 0
    tail := len(height) - 1
    ret := 0
    for {
        if head >= tail {
            break
        }
        tmp := 0
        if height[head] <= height[tail] {
            tmp = height[head] * (tail - head)
            head += 1
        } else {
            tmp = height[tail] * (tail - head)
            tail --
        }
        if tmp > ret {
            ret = tmp
        }
    }
    return ret
}
