package main

import "sort"

// Calculates Median, middle number in a given set
func calcMedian(set []float64) float64 {

	// Sort the input set of numbers
	sort.Float64s(set)

	// Get the middle number of the set by dividing the total number of items in the set by 2
	middleNumber := len(set) / 2

	// Check to see if the total number of items in the set is an odd number. If it is, just return the middle number
	if isOdd(len(set)) {
		return set[middleNumber]
	}

	// If the total number of items in the set is even, return the sum of the two middle numbers divided by 2
	return (set[middleNumber-1] + set[middleNumber]) / 2

}
