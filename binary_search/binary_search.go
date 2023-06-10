package binarysearch

func BinarySearch(arr []int, searchVal int) (int, bool) {
	var lower = 0
	var upper = len(arr) - 1

	for lower <= upper {
		mid := (lower + upper) / 2
		midVal := arr[mid]
		if midVal == searchVal {
			return mid, true
		} else if midVal < searchVal {
			lower = mid + 1
		} else {
			upper = mid - 1
		}
	}
	return 0, false
}
