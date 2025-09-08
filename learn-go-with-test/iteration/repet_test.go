package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
