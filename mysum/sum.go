package mysum

func Sum(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	result := make([]int, len(numbersToSum))
	for i, arr := range numbersToSum {
		result[i] = Sum(arr)
	}
	return result
}
