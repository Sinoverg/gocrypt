package random

func Random(a, b, c int) []int {
	result := make([]int, 0)
	result = append(result, 0)
	for i := 1; i <= 10; i++ {
		result = append(result, ((a*(result[i-1]) + b) % c))
	}
	return result
}
