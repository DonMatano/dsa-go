package insertion_sort

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		searchingValue := arr[i]
		for j := i - 1; j >= 0; j-- {
			if searchingValue < arr[j] {
				temp := arr[j]
				arr[j] = searchingValue
				arr[j+1] = temp
			} else {
				break
			}
		}
	}
	return arr
}
