func strStr(haystack string, needle string) int {
    lenNeedle := len(needle)
    if lenNeedle == 0 {
        return 0
    } else {
        for i := 0; i <= len(haystack) - lenNeedle; i ++ {
            if haystack[i:i+lenNeedle] == needle {
                return i
            }
        } 
    }
    return -1
}