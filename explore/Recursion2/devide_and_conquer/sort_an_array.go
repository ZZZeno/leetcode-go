package chapter1

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	l := len(nums)
        // this shuffle is necessary, can let quick sort worst case's time cost from O(n^2) to O(nlogn)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(l, func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	qSort(nums, 0, l-1)
	return nums
}

func qSort(nums []int, lo, hi int) {
	if lo < hi {
		pivot := partition(nums, lo, hi)
		qSort(nums, lo, pivot-1)
		qSort(nums, pivot+1, hi)
	}
}

func partition(nums []int, lo, hi int) int {
	var pivot = nums[lo]
	var i = lo + 1
	for j := lo + 1; j <= hi; j++ {
		if nums[j] < pivot {
			swap(nums, j, i)
			i++
		}
	}
	swap(nums, i-1, lo)
	return i-1
}

func swap(nums []int, pos1, pos2 int) {
	nums[pos1], nums[pos2] = nums[pos2], nums[pos1]
}
