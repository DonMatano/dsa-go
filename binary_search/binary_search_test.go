package binarysearch_test

import (
	"fmt"
	"testing"

	binarysearch "github.com/DonMatano/dsa-go/binary_search"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		arr        []int
		search     int
		expectIndx int
		expectOk   bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 2, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 0, false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Searching %d in array %v", test.search, test.arr)
		t.Run(testName, func(t *testing.T) {
			ind, ok := binarysearch.BinarySearch(test.arr, test.search)
			if ind != test.expectIndx || ok != test.expectOk {
				t.Errorf("Got index %d and ok %v, instead of index %d and ok %v", ind, ok, test.expectIndx, test.expectOk)
			}
		})
	}
}
