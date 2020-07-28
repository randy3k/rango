package prompt

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}


type Position struct {
    x int
    y int
}
