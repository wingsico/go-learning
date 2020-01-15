package shapes

// Rectangle has width & height props
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter input width & height and get the perimeter
func Perimeter(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}

// Area get the area
func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
