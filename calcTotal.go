package main

func calcTotal(set []float64) float64 {

	var total = 0.0

	for _, v := range set {
		total += v
	}

	return total
}
