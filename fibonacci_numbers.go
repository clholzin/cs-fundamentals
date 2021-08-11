package fundamentals

import (
	"fmt"
)

func fibonacci(n int) {
	var n1, n2, i int
	if n < 1 {
		return
	}

	fmt.Printf("%d ", n1)

	for i = 1; i < n; i++ {
		fmt.Printf("%d ", n2)
		next := n1 + n2
		n1 = n2
		n2 = next
	}
}

func recurseFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return recurseFibonacci(n-1) + recurseFibonacci(n-2)
}
