package main

import "fmt"

func main() {
	var (
		i     int = 0
		sum   int = 0
		LIMIT int = 1000
	)
	for i = 0; i < LIMIT; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
