package searchalgorithms

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// LinearSearchData stores data needed to test the “LinearSearch“ function.
type LinearSearchData[T comparable] struct {
	slice  []T
	value  T
	result int
}

// redirectStdout suppresses output from the test being run.
//
// It returns a closure and error if any.
// You should use defer to close the returned closure.
func redirectStdout() (func(), error) {
	close := func() {}

	// Create a temporory file
	tempFile, err := ioutil.TempFile("", "TestLinearSearch")
	if err != nil {
		return close, err
	}

	// Redirect standard output to the file
	old := os.Stdout
	os.Stdout = os.NewFile(tempFile.Fd(), os.DevNull)

	// Define closure to set Stdout back to old file
	close = func() { os.Stdout = old }

	return close, nil
}

// TestLinearSearch_Int tests the function `LinearSearch` with integer input.
func TestLinearSearch_Int(t *testing.T) {
	// Get a new assertion object for use within this function
	assert := assert.New(t)

	// Create a temporory file and redirect standard output
	close, err := redirectStdout()
	if err != nil {
		log.Panic(err)
	}
	defer close()

	// Get a slice of structs defining the test data
	data := []LinearSearchData[int]{
		// Test for value data that occurs once
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 5},
		// Test for value data that occurs twice
		{[]int{0, 1, 1, 5, 5, 6, 6, 7, 8, 9}, 5, 3},
		// Test for empty slice
		{[]int{}, 0, NotFound},
		// Test for nil slice
		{nil, 10, NotFound},
	}

	for _, datum := range data {
		assert.Equal(datum.result, LinearSearch(datum.slice, datum.value))
	}
}

// TestLinearSearch_String tests the function `LinearSearch` with string input.
func TestLinearSearch_String(t *testing.T) {
	// Get a new assertion object for use within this function
	assert := assert.New(t)

	// Create a temporory file and redirect standard output
	close, err := redirectStdout()
	if err != nil {
		log.Panic(err)
	}
	defer close()

	data := []LinearSearchData[string]{
		// Test for data which occurs once
		{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, "d", 3},
		// Test for data longer that one byte
		{[]string{"cheezaram", "okeke"}, "okeke", 1},
		// Test for data occuring more than once
		{[]string{"okeke", "okeke", "cheezaram", "cheezaram"}, "cheezaram", 2},
		// Test empty slice
		{[]string{}, "cheezaram", NotFound},
		// Test nil slice
		{nil, "okeke", NotFound},
	}

	for _, datum := range data {
		assert.Equal(datum.result, LinearSearch(datum.slice, datum.value))
	}
}
