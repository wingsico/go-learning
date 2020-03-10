package shapes

import (
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

	// 引入 interface， 使用测试辅助函数
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %.2f, want %.2f", got, want)
	// 	}
	// }

	// t.Run("Rectangle Area", func(t *testing.T) {
	// 	shape := Rectangle{10.0, 20.0}
	// 	want := 200.0
	// 	checkArea(t, shape, want)
	// })

	// t.Run("Circle Area", func(t *testing.T) {
	// 	shape := Circle{10.0}
	// 	want := 10.0 * 10.0 * math.Pi
	// 	checkArea(t, shape, want)
	// })

	// 表格驱动测试
	testList := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// {Rectangle{Width: 10, Height: 10}, 100.0},
		// {Circle{R: 10}, 314.1592653589793},
		// {Triangle{B: 10, H: 2}, 10 * 2},
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasArea: 100.0},
		{name: "Circle", shape: Circle{R: 10}, hasArea: 314.1592653589793 * 2},
		{name: "Triangle", shape: Triangle{B: 10, H: 2}, hasArea: 10.0},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasArea
			if tt.hasArea != tt.shape.Area() {
				t.Errorf("%#v want %.2f, got %.2f", tt.shape, want, got)
			}
		})
		// want := tt.want
		// got := tt.shape.Area()
		// if want != got {
		// 	t.Errorf("%#v want %.2f, got %.2f. %#v", tt.shape, want, got, tt.shape)
		// }
	}
}
