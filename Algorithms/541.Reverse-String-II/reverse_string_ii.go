func reverseStr(s string, k int) string {
    byte_s := []byte(s)
    length := len(byte_s)
    for i:=0; i<length; i+=2*k {
        x, y := i, i+k-1
        if y > length-1 {
            y = length - 1
        }
        for ; x<y; x, y = x+1, y-1 {
            byte_s[x], byte_s[y] = byte_s[y], byte_s[x]
        }
    }
    return string(byte_s)
}