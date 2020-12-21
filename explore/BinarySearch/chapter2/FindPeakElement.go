package chapter2

func findPeakElement(nums []int) int {
    ret := 0
    if len(nums) == 1 {
        return 0
    }
    for i := 0; i < len(nums); i ++ {
        if i == 0 {
            if nums[i] > nums[i+1] {
                return i
            }
            continue
        }
        if i == len(nums) - 1 {
            if nums[i] > nums[i-1] {
                return i
            }
            continue
        }
        if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
            return i
        }
    }
    return ret
}
