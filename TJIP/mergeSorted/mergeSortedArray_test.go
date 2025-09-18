package mergeSorted

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	assertEqual := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("merge two arrays", func(t *testing.T) {
		got := merge([]int{1, 2, 3}, 3, []int{2, 5, 6}, 3)
		want := []int{1, 2, 2, 3, 5, 6}
		assertEqual(t, got, want)
	})

	t.Run("empty array merge", func(t *testing.T) {
		got := merge([]int{1}, 1, []int{}, 0)
		want := []int{1}
		assertEqual(t, got, want)
	})
	t.Run("no elements", func(t *testing.T) {
		got := merge([]int{0}, 0, []int{1}, 1)
		want := []int{1}
		assertEqual(t, got, want)
	})
}
