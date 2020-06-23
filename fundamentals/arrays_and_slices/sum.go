package arraysandslices

// Sum : adds up the total of all numbers in an array
func Sum(numbers []int) int {
	sum := 0
	// _ is a blank identifier
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll : gets the Sum for an array of arrays
func SumAll(sliceOfNumbers ...[]int) (sums []int) {
	for _, numbers := range sliceOfNumbers {
		number := Sum(numbers)
		sums = append(sums, number)
	}
	return
}

// SumAllTails : gets the Sum for an array of arrays, but doesn't use the first element
func SumAllTails(sliceOfNumbers ...[]int) (sums []int) {
	for _, numbers := range sliceOfNumbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			number := Sum(numbers[1:])
			sums = append(sums, number)
		}
	}
	return
}
