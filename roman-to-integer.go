func romanToInt(s string) int {
    RtoImap := map[string]int{
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
        "IV": 4, 
        "IX": 9, 
        "XL": 40, 
        "XC": 90,
        "CD":400, 
        "CM":900,
    }
    var num int
    for idx := 0; idx < len(s); idx++ {
        if len(s) >= idx + 2 {
            if val, ok := RtoImap[s[idx:idx+2]]; ok {
                num += val
                idx++
            } else {
                num += RtoImap[string(s[idx])]
            } 
        } else {
            num += RtoImap[string(s[idx])]
        }
    }
    
    return num
}