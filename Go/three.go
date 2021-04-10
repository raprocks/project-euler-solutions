package main

import (
	"fmt"
	"math"
)

func max_prime_factor(num int) int {
	var (
		max_prime int = 1
		sqrt_num  int = int(math.Sqrt(float64(num)))
	)
	for num%2 == 0 {
		num /= 2
	}
	for i := 3; i < sqrt_num; i += 2 {
		if num%i == 0 {
			max_prime = i
			for num%i == 0 {
				num /= i
			}
		}
		if num == 1 {
			break
		}
	}
	return max_prime
}

func main() {
	fmt.Println(max_prime_factor(600851475143))
}
