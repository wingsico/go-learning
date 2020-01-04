package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	excepted := 5

	if excepted != sum {
		t.Errorf("Expected '%d', but got '%d'", excepted, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// output: 5
}
