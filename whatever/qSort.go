package whatever

func partition(items []int, lo, hi int) int {
	i, j := lo, hi+1
	for {


		j -= 1
		for items[j] > items[lo] {
			if j == lo {
				break
			}
			j --
		}
		if i >= j {
			break
		}
		i += 1
		for items[i] < items[lo] {
			if i == hi {
				break
			}
			i ++
		}
		exch(items, i, j)
	}
	exch(items, lo, j)
	return j
}

func exch(items []int, i, j int) {
	temp := items[i]
	items[i] = items[j]
	items[j] = temp
}

func sort(items []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := partition(items, lo, hi)

	sort(items, lo, mid-1)
	sort(items, mid+1, hi)
}

func QuickSort(items []int) {
	//partition(items, 0, len(items)-1)
	sort(items, 0, len(items)-1)
}
