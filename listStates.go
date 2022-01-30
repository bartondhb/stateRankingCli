package main

import (
	"fmt"
	"sort"
)

// Generates an alphabetical list of States and Territories to assist end user with appropriate spelling
func listStates(e *[]Entity) {

	fmt.Printf("List of States and Territories:\n")

	var stateNameList []string

	// Iterate the Entity Slice and add each DisplayName to the stateNameList
	for _, s := range *e {
		stateNameList = append(stateNameList, s.DisplayName)
	}

	// Sort the list
	sort.Strings(stateNameList)

	// Iterate the sorted list and print to Standard Out
	for _, s := range stateNameList {
		fmt.Printf("%s\n", s)
	}

}
