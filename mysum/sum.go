package mysum

func Sum(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// result := make([]int, len(numbersToSum))
	var result []int
	for _, arr := range numbersToSum {
		result = append(result, Sum(arr))
	}
	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	// result := make([]int, len(numbersToSum))
	var result []int
	for _, arr := range numbersToSum {
		if len(arr) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(arr[1:]))
		}
	}
	return result
}
