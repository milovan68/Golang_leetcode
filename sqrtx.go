func mySqrt(x int) int {
    i := 0
    for i * i <= x {
        i ++
    }
    return i - 1
}