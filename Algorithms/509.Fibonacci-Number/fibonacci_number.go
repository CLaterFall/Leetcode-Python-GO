func fib(N int) int {
    f := []int{0, 1}
    for i := 0; i < N; i++ {
        a := f[0]
        b := f[1]
        f[0] = b
        f[1] = a + b
    }
    return f[0]
}