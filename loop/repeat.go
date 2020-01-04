package loop

// Repeat : get s copy five times;
func Repeat(s string) string {
	var repeat string
	for i := 0; i < 5; i++ {
		repeat = repeat + s
	}
	return repeat
}
