func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else {
        y := x
        var reverse int
        for y != 0 {
            remainder := y % 10
            reverse = reverse * 10 + remainder
            y = y / 10
        }
        if reverse == x {
            return true
        } else {
            return false
        }
    }
    
}