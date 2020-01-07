package strings

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestStrings(t *testing.T) {
	testedStr := "6asdbsd6al!~"

	assertCorrectMessage := func(t *testing.T, got, want interface{}) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Compare", func(t *testing.T) {
		got := strings.Compare(testedStr, "bbbb")
		want := -1

		assertCorrectMessage(t, got, want)
	})

	t.Run("Contains", func(t *testing.T) {
		got := strings.Contains(testedStr, "asd")
		want := true

		assertCorrectMessage(t, got, want)
	})

	t.Run("ContainsAny", func(t *testing.T) {
		got := strings.ContainsAny(testedStr, "jjsida")
		want := true

		assertCorrectMessage(t, got, want)
	})

	t.Run("ContainsRune", func(t *testing.T) {
		got := strings.ContainsRune(testedStr, 70)

		want := true

		assertCorrectMessage(t, got, want)
	})
}

func ExampleFields() {
	fmt.Printf("%q", strings.Fields("     a sdasd asd     ttwe sada  "))
	// output: ["a" "sdasd" "asd" "ttwe" "sada"]
}

func ExampleFieldsFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	// output: ["foo1" "bar2" "baz3"]
}

func ExampleSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,.b,c,", ",", -1))
	// Output: ["a," ".b," "c," ""]
}
