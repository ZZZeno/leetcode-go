package problems

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

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	merge(nums1, m, nums2, n)
}
