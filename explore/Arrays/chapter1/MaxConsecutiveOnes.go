package chapter1

//Input: [1,1,0,1,1,1]
//Output: 3
//Explanation: The first two digits or the last three digits are consecutive 1s.
//The maximum number of consecutive 1s is 3.

//Note:
//
//The input array will only contain 0 and 1.
//The length of input array is a positive integer and will not exceed 10,000

func findMaxConsecutiveOnes(nums []int) int {
	l := len(nums)
	max := nums[0]
	tmpMax := 0
	for i := 0; i < l; i ++ {
		tmpMax += nums[i]
		if tmpMax > max {
			max = tmpMax
		}
		if nums[i] == 0 {
			tmpMax = 0
		}
	}
	return max
}

func FindMaxConsecutiveOnes(nums []int) int {
	return findMaxConsecutiveOnes(nums)
}
