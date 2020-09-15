package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Content string
	Time time.Time
}

// Consolidated repeated error checks into a single function
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func readFile(filename string) []string{
	file, err := os.Open(filename)
	checkError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}

// Clean this up later
func getNodeDetails(node string) (string, string){
	lines := readFile("config.txt")
	for _, line := range lines[1:] {
		process := strings.Split(line, " ")
		if process[0] == node {
			return process[1], process[2]
		}
	}
	return "", ""
}

func getDelayParams() (int, int) {
	lines := readFile("config.txt")
	line := strings.Split(lines[0], " ")
	min, _ := strconv.Atoi(line[0])
	max, _ := strconv.Atoi(line[1])
	return min, max
}
