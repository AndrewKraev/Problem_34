package main

import (
	"fmt"
	"math"
)

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Sumfactorial(m int) int {
	var temp int
	var sum int = 0

	for m >= 10 {
		temp = m % 10
		sum = sum + Factorial(temp)
		m = m / 10
	}
	if m < 10 {
		sum = sum + Factorial(m)
	}
	return sum
}

func main() {
	var out_n int

	for i := 1; i*Factorial(9) > int(math.Pow10(i)); i++ {
		out_n = i + 1
	}
	for j := 1; j < int(math.Pow10(out_n)); j++ {
		if j == Sumfactorial(j) {
			fmt.Println(j, "\t", Sumfactorial(j))
		}

	}
}
