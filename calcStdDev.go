package main

import "math"

// Calculates Standard Deviation for a given population (set)
func calcStdDev(set []float64) float64 {

	// First, calculate the mean of the set
	mean := calcMean(set)

	// Establish a variable to hold the running sum of the squared differences between each number in the set and the mean
	var squaredDifferenceSum float64 = 0

	// Iterate the set
	for _, v := range set {

		// For each number in the set, raise the difference between that number and the mean to the power of 2
		squaredDifference := math.Pow(v-mean, 2)

		// Add the squared difference to the running sum of squared differences
		squaredDifferenceSum += squaredDifference

	}

	// Calculate the variance by dividing the total of squared differences by the number of items in the set
	variance := squaredDifferenceSum / float64(len(set))

	// Return the square root of the variance
	return math.Sqrt(variance)

}
