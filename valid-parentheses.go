func isValid(s string) bool {
    var prom []rune
    for _, value := range s {
        if value == '(' || value == '[' || value == '{' {
            prom = append(prom, value)
        } else {
            switch value {
                case ')':
                if len(prom) > 0 && prom[len(prom) - 1] == '('{
                    _, prom = prom[len(prom) - 1], prom[:len(prom) - 1]
                } else {
                    return false
                }
                case ']':
                if len(prom) > 0 && prom[len(prom) - 1] == '['{
                    _, prom = prom[len(prom) - 1], prom[:len(prom) - 1]
                } else {
                    return false
                }
                case '}':
                if len(prom) > 0 && prom[len(prom) - 1] == '{'{
                    _, prom = prom[len(prom) - 1], prom[:len(prom) - 1]
                } else {
                    return false
                }
            }
        }
    }
    if len(prom) == 0{
        return true
    } else {
        return false
    }
}