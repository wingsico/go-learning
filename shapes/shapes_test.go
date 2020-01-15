package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{4, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
