package arrays_and_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum from a collection of any size specified", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sum(numbers)
		want := 28

		if got != want {
			t.Errorf("Expected '%d', got '%d' given %v", want, got, numbers)
		}
	})
}
