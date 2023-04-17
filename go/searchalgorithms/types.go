package searchalgorithms

import (
	"io/ioutil"
	"log"
	"os"
)

// NotFound is a constant for the return value of the search algorithm
const NotFound = -1
const logFile = "/tmp/searchalgorithms.log"
const logFilePerm = 0644

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

// setUpLogFile sets the logging file
func setUpLogFile() (func(), error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, logFilePerm)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)
	close := func() { file.Close() }
	return close, nil
}
