package searchalgorithms

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"
)

// BinarySearch finds the index of the value in the slice.
//
// It takes in an Ordered slice and a value to find and returns the index of the value
// or “NotFound“ if the value is not found in the slice.
func BinarySearch[T constraints.Ordered](slice []T, value T) (int, error) {
	// Setup logging file
	close, err := setUpLogFile()
	if err != nil {
		return 0, fmt.Errorf("BinarySearch failed toopen log file: %w", err)
	}
	defer close()

	log.Println("In BinarySearch")

	var left, right = 0, len(slice) - 1

	for left <= right {
		pivot := (left + right) / 2
		log.Printf("Value checked at index [%d] = [%v]\n", pivot, slice[pivot])

		switch {
		case slice[pivot] > value:
			right = pivot - 1
		case slice[pivot] < value:
			left = pivot + 1
		default:
			log.Printf("Found [%v] at index [%d]\n\n", value, pivot)
			return pivot, nil
		}
	}

	log.Printf("Value [%v] not found in slice\n\n", value)
	return NotFound, nil
}
