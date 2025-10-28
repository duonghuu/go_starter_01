package utils

func Factorial(n int) int {
	result := 1
	i := n
	for i > 1 {
		result *= i
		i--
	}

	return result
}
