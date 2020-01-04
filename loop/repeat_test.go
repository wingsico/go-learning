package loop

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	excepted := "aaaaa"

	if got != excepted {
		t.Errorf("excepted %q, but got %q", excepted, got)
	}
}
