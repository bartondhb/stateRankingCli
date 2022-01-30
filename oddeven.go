package main

func isEven(d int) bool {
	// Determine if a number is even by checking for a remainder. If there is none, the number is even
	return d%2 == 0
}

func isOdd(d int) bool {
	// Determine if a number is odd by checking that it's not even :-)
	return !isEven(d)
}
