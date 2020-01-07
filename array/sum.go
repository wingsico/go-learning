package array

// Sum get a sum of a array
func Sum(a []int) (sum int) {
	sum = 0

	for _, n := range a {
		sum += n
	}
	return sum
}
