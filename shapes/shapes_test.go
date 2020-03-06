package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	// t.Run("rectangle", func(t *testing.T) {
	// 	got := Rectangle{4, 10.0}.Area()
	// 	want := 40.0

	// 	if got != want {
	// 		t.Errorf("got %v, want %v", got, want)
	// 	}
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	got := Circle{10}.Area()
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("got %v, want %v", got, want)
	// 	}
	// })

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("Rectangle Area", func(t *testing.T) {
		shape := Rectangle{10.0, 20.0}
		want := 200.0
		checkArea(t, shape, want)
	})

	t.Run("Circle Area", func(t *testing.T) {
		shape := Circle{10.0}
		want := 10.0 * 10.0 * math.Pi
		checkArea(t, shape, want)
	})
}
