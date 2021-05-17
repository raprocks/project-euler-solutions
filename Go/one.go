package solutions

// Solution function gives out the sum of the numbers that are multiples divisible by 3 or 5 till the LIMIT.
func SolutionOne(LIMIT int) int {
	var (
		i   int = 0 // counter
		sum int = 0 // total of multiples
	)
	for i = 0; i < LIMIT; i++ { // for loop from 0 to LIMIT-1
		if i%3 == 0 || i%5 == 0 { // if conditional to check if i is multiple of 3 or 5
			sum += i // add to sum
		}
	}
	return sum // return sum
}
