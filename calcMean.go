package main

// Calculates Mean (Simple Average)
func calcMean(set []float64) float64 {

	// Calculate the total of all numbers in the set
	var total = calcTotal(set)

	// Divide the total by the number of items in the set
	return total / float64(len(set))

}
