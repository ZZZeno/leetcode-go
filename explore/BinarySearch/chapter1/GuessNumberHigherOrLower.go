package chapter1

/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    if n == 1 {
        return 1
    }
    lo := 1
    hi := n
    var mid int
    for lo < hi {
        mid = (lo+hi) / 2
        if hi - lo == 1 {
            if guess(hi) == 0 {
                return hi
            }
            return lo
        }
        if guess(mid) == 0 {
            return mid
        }
        if guess(mid) < 0 {
            hi = mid
        } else {
            lo = mid
        }
    }
    return mid
}
