package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var summaryFlag = flag.Bool("summary", false, "Prints a summary of population data for all U.S. States and Territories")
	var stateFlag = flag.String("state", "", "Prints population of the specified state.\n\tExample:\n\t-state california\n\tPrints population of California\n\n\t-state list\n\tPrints an alphabetical list of all U.S. States and Territories\n\n\t-state all\n\tPrints the population of all states in the list.\n")
	flag.Parse()

	sourceUrl := "https://en.wikipedia.org/wiki/List_of_U.S._states_and_territories_by_population"

	data, err := getStateAndTerritoryData(sourceUrl)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	if *summaryFlag {
		summarize(&data)

	} else if *stateFlag != "" {

		switch strings.ToLower(*stateFlag) {
		case "list":
			listStates(&data)
		case "all":
			for _, v := range data {

				err := populationByState(&data, v.Name)

				if err != nil {
					fmt.Printf("%s\n", err.Error())
					os.Exit(1)
				}

			}
		default:
			err := populationByState(&data, strings.ToLower(*stateFlag))

			if err != nil {
				fmt.Printf("%s\n", err.Error())
				os.Exit(1)
			}
		}

		return

	} else {
		flag.Usage()
	}

}
