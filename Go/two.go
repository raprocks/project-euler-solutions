package main

import "fmt"

func fib_sum_even(limit int) int {
	var (
		sum int = 0
		a   int = 1
		b   int = 2
	)

	for b < limit {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}

func main() {
	sum := fib_sum_even(4000000)
	fmt.Println(sum)
}
