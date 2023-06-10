package binarysearch

// Search the index of a search value in given array. If there value is not available in the array
// it returns ok as false.
func BinarySearch(arr []int, searchVal int) (indx int, ok bool) {
	// The lower and upper limits are the first index[0] and last index of the array
	var lower = 0
	// We minus to get the index
	var upper = len(arr) - 1

	//  As long the lower limit value has not overlap the upper limit
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
