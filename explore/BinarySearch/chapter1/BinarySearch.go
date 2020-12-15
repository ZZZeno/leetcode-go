package chapter1

func search(nums []int, target int) int {
    if len(nums) == 1 {
        if nums[0] == target {
            return 0
        }
        return -1
    }
    lo := 0
    hi := len(nums) - 1
    var mid int
    
    for lo < hi {
        mid = (lo+hi) / 2
        if nums[mid] == target {
            return mid
        }
        if lo == mid {
            if nums[hi] == target {
                return hi
            }
            return -1
        }
        if nums[mid] < target {
            lo = mid
        } else {
            hi = mid
        }
    }
    return -1
}
