package chapter2
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    if n == 1 {
        return 1
    }
    left, right := 1, n
    for left < right {
        mid := left + (right-left)/2
        if mid == left {
            if isBadVersion(left) {
                return left
            }
            return right
        }
        if isBadVersion(mid) {
            if (mid-1 > 0 && !isBadVersion(mid-1)) || mid == 1 {
                return mid
            }
            right = mid
            continue
        }
        left = mid
    }
    return left
}
