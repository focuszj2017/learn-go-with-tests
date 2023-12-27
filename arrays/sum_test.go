package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("test slices equal", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		numbers2 := []int{1, 2, 3}

		isEqual := slices.Equal(numbers, numbers2)

		if !isEqual {
			t.Errorf("not equal %v %v",
				numbers, numbers2)
		}
	})

}
