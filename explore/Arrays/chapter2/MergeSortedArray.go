package chapter2

//	Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//	
//	Note:
//	
//	The number of elements initialized in nums1 and nums2 are m and n respectively.
//	You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
//	Example:
//	
//	Input:
//	nums1 = [1,2,3,0,0,0], m = 3
//	nums2 = [2,5,6],       n = 3
//	
//	Output: [1,2,2,3,5,6]
//	
//	
//	Constraints:
//	
//	-10^9 <= nums1[i], nums2[i] <= 10^9
//	nums1.length == m + n
//	nums2.length == n


func merge(nums1 []int, m int, nums2 []int, n int)  {
	tmp := make([]int, m)
	copy(tmp, nums1)
	i, j := 0, 0
	for ; i < m && j < n; {
		if tmp[i] < nums2[j] {
			nums1[i+j] = tmp[i]
			i += 1
		} else {
			nums1[i+j] = nums2[j]
			j += 1
		}
	}
	if i == m {
		for k := j; i+k < n+m; k ++ {
			nums1[i+k] = nums2[k]
		}
	}
	if j == n {
		for k := i; j+k < n+m; k ++ {
			nums1[j+k] = tmp[k]
		}
	}
}
