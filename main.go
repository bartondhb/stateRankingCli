package main

import (
	"flag"
	"fmt"
	"log"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"sort"
	"strings"
)

func main () {

	sourceUrl := "https://en.wikipedia.org/wiki/List_of_U.S._states_and_territories_by_population"

	data, err := getStateAndTerritoryData(sourceUrl)

	if err != nil {
		log.Fatal(err)
	}

	var summaryFlag = flag.Bool("summary", false, "summarize data for all states and territories")
	var stateFlag = flag.String("state", "", "retrieve population for the named state")
	flag.Parse()

	if *summaryFlag {
		err := summarize(&data)

		if err != nil {
			log.Fatal(err)
		}

	} else if *stateFlag != "" {

		if *stateFlag == "list" {
			listStates(&data)
			return
		}
		err := byState(&data, strings.ToLower(*stateFlag))

		if err != nil {
			log.Fatal(err)
		}

	} else {
		flag.Usage()
	}

}


func summarize(e *[]Entity) error {

	//var totalPopulation int64 = 0
	var statePopulations []float64

	for _,s := range *e {
		//totalPopulation += s.Population
		statePopulations = append(statePopulations, float64(s.Population))
	}

	p := message.NewPrinter(language.English)

	p.Printf("As of the 2020 Census (all values rounded to nearest whole number):\n\nThe total population of the United States is %d.\n", int64(math.Round(calcTotal(statePopulations))))

	p.Printf("The mean population of the United States is %d.\n", int64(math.Round(calcMean(statePopulations))))
	p.Printf("The median population of the United States is %d.\n", int64(math.Round(calcMedian(statePopulations))))

	return nil
}

func calcTotal (set []float64) float64 {

	var total = 0.0

	for _,v := range set {
		total += v
	}

	return total
}

func isEven(d int) bool {
	// Determine if a number is even by checking for a remainder. If there is none, the number is even
	return d % 2 == 0
}

func isOdd(d int) bool {
	// Determine if a number is odd by checking that it's not even :)
	return !isEven(d)
}

func calcMean(set []float64) float64 {

	// Calculate the total of all numbers in the set
	var total = calcTotal(set)

	// Divide the total by the number of items in the set
	return total / float64(len(set))

}

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

func calcStdDev(set []float64)  {

	mean := calcMean(set)


}

func listStates(e *[]Entity) {

	fmt.Printf("List of States and Territories:\n" )

	var stateNameList []string

	for _,s := range *e {
		stateNameList = append(stateNameList, s.DisplayName)
	}

	sort.Strings(stateNameList)

	for _,s := range stateNameList {
		fmt.Printf("%s\n", s)
	}

}

func byState(e *[]Entity, query string) error {

	for _, s := range *e {
		if s.Name == query {

			p := message.NewPrinter(language.English)

			_, err := p.Printf("As of the 2020 Census, the population of %s is %d.\n", s.DisplayName, s.Population)

			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("no state or territory found with name %s\n", query)
}

