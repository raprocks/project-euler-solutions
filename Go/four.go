package main

import "fmt"

func palindrome(num int) bool {
	var (
		num_copy int = num
		rev      int = 0
	)
	for num != 0 {
		rev = (rev * 10) + (num % 10)
		num /= 10
	}
	return (rev == num_copy)
}

func main() {
	var (
		largest_palindrome int = 0
		prod               int = 0
	)
	for i := 999; i > 100; i-- {
		for j := i; i > 100; j-- {
			prod = i * j
			if (i * j) < largest_palindrome {
				break
			}
			if palindrome(prod) {
				if largest_palindrome < prod {
					largest_palindrome = prod
				}
			}
		}
	}
	fmt.Println(largest_palindrome)
}
