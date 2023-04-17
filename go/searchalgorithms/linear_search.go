package searchalgorithms

import (
	"fmt"
	"log"
)

// LinearSearch finds the index of a value in a slice. It operates on comparable
// data types i.e. int, string, float64, etc.
//
// It takes in a slice and a value and returns the index of the first occurence
// of the value if it is found, otherwise it returns “NotFound“.
func LinearSearch[T comparable](slice []T, value T) (int, error) {
	// Set up a log file
	close, err := setUpLogFile()
	if err != nil {
		return 0, fmt.Errorf("LinearSearch failed to open file: %w", err)
	}
	defer close()

	log.Println("In LinearSearch")

	for i, item := range slice {
		log.Printf("Value checked at index [%v] = [%v]\n", i, item)
		if item == value {
			log.Printf("Value found at index [%v] = [%v]\n\n", i, item)
			return i, nil
		}
	}

	log.Printf("Value [%v] not found in slice\n\n", value)
	return NotFound, nil
}
