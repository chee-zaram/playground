package searchalgorithms

import (
	"fmt"
	"log"
)

// InterpolationSearch is a search algorithm that finds a value in a sorted slice of values.
//
// It takes a sorted slice of integers and the value to search for
// and returns the index of the first occurence of the value in the slice
// or “NotFound“ if the value is not found.
func InterpolationSearch(slice []int, value int) (int, error) {
	// Set up logging file
	close, err := setUpLogFile()
	if err != nil {
		return 0, fmt.Errorf("InterpolationSearch failed to open file: %w", err)
	}
	defer close()

	log.Printf("Performing Interpolation Search on slice %v\n", slice)

	blockStartIndex := 0
	blockEndIndex := len(slice) - 1

	// Use interpolation formula to find the index of the value
	for blockStartIndex <= blockEndIndex && slice[blockStartIndex] <= value && slice[blockEndIndex] >= value {
		pos := blockStartIndex + ((blockEndIndex-blockStartIndex)*(value-slice[blockStartIndex]))/(slice[blockEndIndex]-slice[blockStartIndex])
		log.Printf("Value at index [%d] = [%d]\n", pos, slice[pos])

		switch {
		case slice[pos] > value:
			blockEndIndex = pos - 1
		case slice[pos] < value:
			blockStartIndex = pos + 1
		default:
			// Make sure it is the first occurence
			for pos > 0 && slice[pos-1] == value {
				pos--
			}
			log.Printf("Value [%v] found at index [%v]\n\n", value, pos)
			return pos, nil
		}
	}

	// Value was not found
	log.Printf("Value [%v] not found in slice\n\n", value)
	return NotFound, nil
}
