package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// getStateAndTerritoryData retrieves the population data from Wikipedia given the correct URL and parse the HTML as XML
// to extract and create a slice of Entity objects containing data for each State and Territory.
func getStateAndTerritoryData(url string) ([]Entity, error) {

	// Retrieve the data from the supplied url
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	// Handle non-200 response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error retrieving source data: %s", resp.Status)
	}

	// Read the response
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// Deserialize the HTML to a custom HTML struct so that we can handle the resulting HTML as standard XML
	var html HTML

	err = xml.Unmarshal(bodyBytes, &html)

	if err != nil {
		return nil, err
	}

	// Not all <div> elements in the source body have an ID attribute. This helper function returns the <div>
	// closest to the table data I need so that I don't have to traverse the entire XML tree by key.
	mydiv := html.Body.Div.getById("mw-content-text")

	// Create an array to hold data for each State or Territory
	var stateData []Entity

	// Iterates through the Rows in the first Table of the first inner <div> of  the "mw-content-text" <div>
	for _, v := range mydiv.Div[0].Table[0].Body.Row {

		// Most states and territories in the table are handled by this code...
		if len(v.Data) != 0 {

			// If the third data element in the row contains an HTML Anchor Tag...
			if len(v.Data[2].A) != 0 {

				// Get the name of the state from the Value of the first Anchor in the slice
				stateName := v.Data[2].A[0].Value
				// Get the state's 2020 census population from the 4th data element's value, remove string formatting (newlines and commas), and convert it to an integer
				statePop, err := strconv.Atoi(strings.Trim(strings.ReplaceAll(v.Data[3].Value, ",", ""), "\n"))

				// Handle any string to integer conversion errors
				if err != nil {
					return nil, fmt.Errorf("error converting population value to integer: %e", err)
				}

				// Append the data for the state or territory to the stateData slice (and cast integer to int64)
				stateData = append(stateData, Entity{Name: strings.ToLower(stateName), DisplayName: stateName, Population: int64(statePop)})

			}

			// This code does everything the block above does, but it handles a slight formatting difference in the table for Northern Mariana Islands and District of Columbia
			if len(v.Data[2].Span) != 0 {

				// Gets the stateName value from the Anchor inside the "nowrap" span
				if len(v.Data[2].A) == 0 && v.Data[2].Span[0].Class == "nowrap" {

					stateName := v.Data[2].Span[0].A[0].Value
					statePop, err := strconv.Atoi(strings.Trim(strings.ReplaceAll(v.Data[3].Value, ",", ""), "\n"))

					if err != nil {
						return nil, fmt.Errorf("error converting population value to integer: %e", err)
					}

					stateData = append(stateData, Entity{Name: strings.ToLower(stateName), DisplayName: stateName, Population: int64(statePop)})
				}
			}

		}

	}

	return stateData, nil

}
