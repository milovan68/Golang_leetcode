func reverse(x int) int {

    y := x
    if y < 0 {
        y = -y
    }
    var reverse int
    for y != 0 {
        remainder := y % 10
        reverse = reverse * 10 + remainder
        y = y / 10
    }
    if reverse > ((1 << 31) - 1) || reverse < -(1 << 31){
        return 0
    } else {
        if x < 0 {
            return -reverse
        } else {
            return reverse
        }
    }
    
}
