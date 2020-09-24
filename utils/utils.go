package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Message struct {
	Content string
}

// Consolidated repeated error checks into a single function
func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// Read the given text file and return an array with each element being a line from the file
func ReadFile(filename string) []string{
	file, err := os.Open(filename)
	CheckError(err)

	// Create scanner object and textlines array
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	// Loop through file lines, appending to textlines
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	return txtlines
}

// Given a node ID, return the IP and PORT of that node
func GetNodeDetails(node string) (string, string, error){

	// Loop through config file lines, skipping line 1 since it contains delay params
	lines := ReadFile("config.txt")
	for _, line := range lines[1:] {

		// For each line, if the first string matches the ID, return the IP and PORT
		process := strings.Split(line, " ")
		if process[0] == node {
			return process[1], process[2], nil
		}
	}

	// If there's no node with the given ID, return empty strings and an error
	return "", "", errors.New("no node with that id")
}

// Returns delay parameters specified in config.txt
func GetDelayParams() (int, int) {

	// Get lines from the config file
	lines := ReadFile("config.txt")

	// Split the first line into two elements: First is min delay, second is max.
	line := strings.Split(lines[0], " ")
	min, _ := strconv.Atoi(line[0])
	max, _ := strconv.Atoi(line[1])
	return min, max
}
