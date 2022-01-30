package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
)

// Summarize population data for the entire United States
func summarize(e *[]Entity) {

	var statePopulations []float64

	for _, s := range *e {
		statePopulations = append(statePopulations, float64(s.Population))
	}

	p := message.NewPrinter(language.English)

	p.Printf("As of the 2020 U.S. Census (all values rounded to nearest whole number):\n\n")

	p.Printf("The TOTAL population of the United States is %d.\n", int64(math.Round(calcTotal(statePopulations))))
	p.Printf("The MEAN population of States and Territories within the United States is %d.\n", int64(math.Round(calcMean(statePopulations))))
	p.Printf("The MEDIAN population of States and Territories within the United States is %d.\n", int64(math.Round(calcMedian(statePopulations))))
	p.Printf("The STANDARD DEVIATION of States and Territories within the United States is %d.\n", int64(math.Round(calcStdDev(statePopulations))))

}
