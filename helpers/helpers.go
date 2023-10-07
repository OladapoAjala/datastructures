package helpers

import "math"

func IsPrime(num int32) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%int32(i) == 0 {
			return false
		}
	}
	return true
}

func NextPrime(input int32) int32 {
	input++
	for !IsPrime(input) {
		input++
	}
	return input
}
