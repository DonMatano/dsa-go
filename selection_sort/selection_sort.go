package selectionsort

func SelectionSort (arr[]int) []int {
  for i := 0; i < len(arr); i++ {
    lowestValueIndex := i
    for j := i+1; j < len(arr); j++ {
      if (arr[j] < arr[lowestValueIndex]){
        lowestValueIndex = j
      }
    }
    temp := arr[i]
    arr[i] = arr[lowestValueIndex]
    arr[lowestValueIndex] = temp
  }
  return arr
}
