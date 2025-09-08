package arrayAndSlices

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	// clean up code
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {
	// lenOfNumber := len(numberToSum)
	// sums := make([]int, lenOfNumber)
	// for i, number := range numberToSum {
	// 	sums[i] = Sum(number)
	// }
	// return sums
	return nil
}

func SumAllTails(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
