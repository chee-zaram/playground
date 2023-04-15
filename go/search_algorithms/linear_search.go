package searchalgorithms

import (
	"fmt"
	"log"
	"os"
)

// LinearSearch finds the index of a value in a slice. It operates on comparable
// data types i.e. int, string, float64, etc.
//
// It takes in a slice and a value and returns the index of the first occurence
// of the value if it is found, otherwise it returns “NotFound“.
func LinearSearch[T comparable](slice []T, value T) (int, error) {
	// Set up a log file
	file, err := os.OpenFile("/tmp/searchalgorithms.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, fmt.Errorf("LinearSearch: failed to open file: %w", err)
	}

	log.SetOutput(file)
	log.Println("In LinearSearch")

	for i, item := range slice {
		log.Printf("Value checked at index [%v] = [%d]\n", item, i)
		if item == value {
			log.Printf("Value found at index [%v] = [%d]\n\n", item, i)
			return i, nil
		}
	}

	log.Printf("Value [%v] not found in slice\n\n", value)
	file.Close()
	return NotFound, nil
}
