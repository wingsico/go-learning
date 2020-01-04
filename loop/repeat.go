package loop

// Repeat : got a string repeat by $count times;
func Repeat(s string, count int) string {
	var repeat string
	for i := 0; i < count; i++ {
		repeat += s
	}
	return repeat
}
