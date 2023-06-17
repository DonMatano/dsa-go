package datastructures_test

import (
	"testing"

	datastructures "github.com/DonMatano/dsa-go/data_structures"
)

func TestQueue(t *testing.T) {
	t.Run("adding to queue works", func(t *testing.T) {
		queue := datastructures.NewQueue[int]()
		expectedVal := 5
		queue.Enqueue(5)
		val, _ := queue.Read()
		if val != expectedVal {
			t.Errorf("Expected %v, got %v", expectedVal, val)
		}
	})

	t.Run("reading from empty queue return ok:false", func(t *testing.T) {
		queue := datastructures.NewQueue[int]()
		expectedVal := false
		_, ok := queue.Read()
		if ok != expectedVal {
			t.Errorf("Expected %v, got %v", expectedVal, ok)
		}
	})

	t.Run("equeueing and dequeueing works", func(t *testing.T) {
		queue := datastructures.NewQueue[int]()
		for i := range [10]int{} {
			queue.Enqueue(i)
		}
		for i := range [10]int{} {
			v, _ := queue.Dequeue()
			if i != v {
				t.Errorf("Expected %v, got %v", i, v)
			}
		}
		_, ok := queue.Read()
		if ok != false {
			t.Errorf("Expected %v, got %v", false, ok)
		}
	})
}
