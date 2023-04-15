package searchalgorithms

import "fmt"

// LinearSearch finds the index of a value in a slice. It operates on comparable
// data types i.e. int, string, float64, etc.
//
// It takes in a slice and a value and returns the index of the first occurence
// of the value if it is found, otherwise it returns “NotFound“.
func LinearSearch[T comparable](slice []T, value T) int {
	for i, item := range slice {
		fmt.Printf("Value checked at index [%v] = [%d]\n", item, i)
		if item == value {
			fmt.Printf("Value found at index [%v] = [%d]\n\n", item, i)
			return i
		}
	}

	fmt.Printf("Value [%v] not found in slice\n\n", value)
	return NotFound
}
