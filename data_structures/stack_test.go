package datastructures_test

import (
	"testing"

	datastructures "github.com/DonMatano/dsa-go/data_structures"
)

func TestStack(t *testing.T) {
	t.Run("adding to stack works", func(t *testing.T) {
		stack := datastructures.NewStack[int]()
		expectedVal := 5
		stack.Stack(5)
		val, _ := stack.Read()
		if val != expectedVal {
			t.Errorf("Expected %v, got %v", expectedVal, val)
		}
	})

	t.Run("reading from empty stack return ok:false", func(t *testing.T) {
		stack := datastructures.NewStack[int]()
		expectedVal := false
		_, ok := stack.Read()
		if ok != expectedVal {
			t.Errorf("Expected %v, got %v", expectedVal, ok)
		}
	})

	t.Run("stacking and destacking works", func(t *testing.T) {
		stack := datastructures.NewStack[int]()
		for i := range [10]int{} {
			stack.Stack(i)
		}
		for _, val := range []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} {
			v, _ := stack.Destack()
			if val != v {
				t.Errorf("Expected %v, got %v", val, v)
			}
		}
		_, ok := stack.Read()
		if ok != false {
			t.Errorf("Expected %v, got %v", false, ok)
		}
	})
}
