package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat char based on user provided argument", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		if repeated != expected {
			t.Errorf("Expected %q, got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
