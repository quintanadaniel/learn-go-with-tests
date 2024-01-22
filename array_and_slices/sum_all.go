package arrayandslices

func SumAll(numbersToSum ...[]int) []int {
	//lenghtOfNumbers := len(numbersToSum)
	//sums := make([]int, lenghtOfNumbers)

	//for i, numbers := range numbersToSum {
	//	sums[i] = Sum(numbers)
	//}

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numberOfSumTails ...[]int) []int {
	var sums []int

	for _, numbers := range numberOfSumTails {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
