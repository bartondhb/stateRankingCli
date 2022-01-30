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

	p.Printf("The total population of the United States is %d.\n", int64(math.Round(calcTotal(statePopulations))))
	p.Printf("The mean population of the United States is %d.\n", int64(math.Round(calcMean(statePopulations))))
	p.Printf("The median population of the United States is %d.\n", int64(math.Round(calcMedian(statePopulations))))
	p.Printf("The standard deviation of the population of the United States is %d.\n", int64(math.Round(calcStdDev(statePopulations))))

}
