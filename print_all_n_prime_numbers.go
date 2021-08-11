package fundamentals

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printPrime(n int) {
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
