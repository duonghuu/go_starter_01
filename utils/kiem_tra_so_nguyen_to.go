package utils

func PrimeNumber(n int) bool {
	if n < 1 {
		return false
	}
	i := 2
	for i < (n - 1) {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}
