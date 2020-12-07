package problems

func isUgly(num int) bool {
    if num == 1 {
        return true
    }
    if num == 0 {
        return false
    }
    div := num
    mod := num
    
    mod = div % 5
    for mod == 0 {
        div = div / 5
        mod = div % 5
    }

    mod = div % 3
    for mod == 0 {
        div = div / 3
        mod = div % 3
    }
    
    mod = div % 2
    for mod == 0 {
        div = div / 2
        mod = div % 2
    }
    
    return div == 1
}
