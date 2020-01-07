package array

// Sum get a sum of a array
func Sum(a []int) (sum int) {
	sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}
