package searchalgorithms

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/constraints"
)

// ExponentialSearch increases exponentially with each iteration until it finds a
// block where the target value can be found. It then performs a binary search on
// that block to find the target value.
//
// The function takes in a slice of any Ordered type and a target value. It returns
// the index of the first occurrence of the target value or NotFound if target is not found,
// and error (if any).
func ExponentialSearch[T constraints.Ordered](slice []T, value T) (int, error) {
	// Set up logging file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, logFilePerm)
	if err != nil {
		return 0, fmt.Errorf("ExponentialSearch could not open log file: %w", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Printf("Performing Exponential Search on slice %v for target %v\n", slice, value)

	size := len(slice)
	var upperBound, lowerBound = 1, 0

	// Find the suitable block where value could be
	for upperBound < size && slice[upperBound] < value {
		log.Printf("Value at index [%v] = [%v]\n", upperBound, slice[upperBound])
		lowerBound = upperBound
		upperBound *= 2
	}

	// Make sure we're not out of bounds
	if upperBound >= size {
		upperBound = size - 1
	}

	log.Printf("Value [%v] should be in block[%v,%v]\n", value, lowerBound, upperBound)
	// Perform binary search on the block to find an occurrence
	reval, err := BinarySearch(slice[lowerBound:upperBound+1], value)
	if err != nil {
		return 0, fmt.Errorf("ExponentialSearch could not perform binary search: %w", err)
	}

	// Value was found. Now, make sure it's the first occurrence
	if reval != NotFound {
		pos := reval + lowerBound
		for pos > 0 && slice[pos-1] == value {
			pos--
		}
		log.Printf("In ExponentialSearch, value [%v] found at index [%v]\n\n", value, pos)
		return pos, nil
	}

	log.Printf("Value [%v] not found in slice\n\n", value)
	return NotFound, nil
}
