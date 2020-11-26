package chapter5

func replaceElements(arr []int) []int {
    if len(arr) == 1 {
        arr[0] = -1
        return arr
    }
    max := arr[len(arr)-1]
    arr[len(arr)-1] = -1
    for i := len(arr)-2; i >= 0; i -- {
        temp := arr[i]
        arr[i] = max
        if temp > max {
            max = temp
        }
    }
    return arr
}
