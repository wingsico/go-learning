package array

// Sum get a sum of a array
func Sum(a []int) (sum int) {
	sum = 0

	for _, n := range a {
		sum += n
	}
	return
}

// SumAll get a sum slice of some slices
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumTail get a tail sum slice of some slices
func SumTail(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return
}
