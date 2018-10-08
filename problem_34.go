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

	if m < 10 {
		sum = sum + Factorial(m)
	}

	for m > 10 {
		temp = m % 10
		sum = sum + Factorial(temp)
		m = m / 10
	}
	return sum
}

func main() {
	var out_n int

	for i := 1; i*Factorial(9) > int(math.Pow10(i)); i++ {
		//		fmt.Println(i * Factorial(9))
		//		fmt.Println(int(math.Pow10(i)) - 1)
		out_n = i + 1
	}
	fmt.Println("Предел для проверки: ", int(math.Pow10(out_n))-1)
	fmt.Println("Возможная сумма цифр: ", out_n*Factorial(9))

	for j := 1; j < int(math.Pow10(out_n)); j++ {
		if j == Sumfactorial(j) {
			fmt.Println(j, " ", Sumfactorial(j))
		}

	}
}
