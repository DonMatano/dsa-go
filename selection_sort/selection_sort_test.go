package selectionsort_test

import (
	"fmt"
	"testing"

	selectionsort "github.com/DonMatano/dsa-go/selection_sort"
)

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSelectionSort(t *testing.T) {
  var tests = []struct {
    arr, expect []int
  }{
		{[]int{5, 7, 6, 1, 4, 2, 3}, []int{1, 2, 3, 4, 5, 6, 7}},
  }
	for _, test := range tests {
		testName := fmt.Sprintf("Sorting arr, %v", test.arr)
		t.Run(testName, func(t *testing.T) {
			sortedArr := selectionsort.SelectionSort(test.arr)
			if !testEq(sortedArr, test.expect) {
				t.Errorf("Expected %v, got %v", test.expect, sortedArr)
			}
		})
	}
}
