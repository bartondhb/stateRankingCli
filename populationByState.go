package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Show the population of a specified State or Territory
func populationByState(e *[]Entity, query string) error {

	for _, s := range *e {
		if s.Name == query {

			p := message.NewPrinter(language.English)

			p.Printf("As of the 2020 U.S. Census, the population of %s is %d.\n", s.DisplayName, s.Population)

			return nil
		}
	}

	return fmt.Errorf("no state or territory found with name %s\n", query)
}
