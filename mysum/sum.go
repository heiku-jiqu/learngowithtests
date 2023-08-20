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
