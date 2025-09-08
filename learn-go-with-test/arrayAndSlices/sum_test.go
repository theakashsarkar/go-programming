package arrayAndSlices

import (
	"reflect"
	"testing"
)

// func TestSum(t *testing.T) {

// 	t.Run("collection of 5 number", func(t *testing.T) {
// 		numbers := []int{1, 2, 3, 4, 5}
// 		got := Sum(numbers)
// 		want := 15

// 		if got != want {
// 			t.Errorf("got %d want %d given, %v", got, want, numbers)
// 		}
// 	})
// 	t.Run("collection of any size", func(t *testing.T) {
// 		numbers := []int{1, 2, 3}
// 		got := Sum(numbers)
// 		want := 6
// 		if got != want {
// 			t.Errorf("got %d want %d given, %v", got, want, numbers)
// 		}
// 	})

// 	// got := TestSumAll([]int{1, 3}, []int{0, 9})
// 	// want := []int{3, 9}

// 	// if !reflect.DeepEqual(got, want) {
// 	// 	t.Errorf("got %v want %v", got, want)
// 	// }

// 	got := SumAll([]int{1, 2}, []int{0, 9})
// 	want := "bob"

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
