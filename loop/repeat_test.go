package loop

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	excepted := "aaaaaaa"

	if got != excepted {
		t.Errorf("excepted %q, but got %q", excepted, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("test", 3)
	fmt.Println(res)
	// output: testtesttest
}
