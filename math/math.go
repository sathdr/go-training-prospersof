package math

import "fmt"

// Power calculate power
func Power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

// Prime is the function for printing prime numbers
func Prime(n int) {
	for i := 1; i <= n; i++ {
		if IsPrime(i) {
			fmt.Print(i)
			fmt.Print(",")
		}
	}
	fmt.Println()
}

// IsPrime is the function for checking prime number
func IsPrime(n int) bool {
	count := 0
	for j := 1; j <= n; j++ {
		if n%j == 0 {
			count++
		}
	}
	return count == 2
}
