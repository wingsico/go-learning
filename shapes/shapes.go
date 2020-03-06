package shapes

import "math"

// Rectangle has width & height props
type Rectangle struct {
	Width  float64
	Height float64
}

// Shape : a interface with Area method
type Shape interface {
	Area() float64
}

// Area : Compute the rect's area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has radius
type Circle struct {
	R float64
}

// Area : Compute the circle's area
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

// Perimeter input width & height and get the perimeter
func Perimeter(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}

// Area get the area
func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
