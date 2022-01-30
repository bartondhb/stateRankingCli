package main

// Determine if a number is even by checking for a remainder. If there is none, the number is even
func isEven(d int) bool {
	return d%2 == 0
}

// Determine if a number is odd by checking that it's not even :-)
func isOdd(d int) bool {
	return !isEven(d)
}
