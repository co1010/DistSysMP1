package main

import (
	"MP1/utils"
	"bufio"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"strings"
	"time"
)

func main() {

	// Create new parser object
	parser := argparse.NewParser("Process", "Saves which process is running")

	// Create string flag
	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "Input Process ID"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	// Start the TCP listen server for specified ID
	go startServer(s)

	// Read input from the command line.
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		input := strings.Split(text, " ")
		content := strings.Join(input[2:], " ")
		content = strings.TrimRight(content, "\r\n")

		// If the user enters the send command, call unicast_send as a goroutine
		if input[0] == "send" {
			message := utils.Message{content, time.Now()}
			go unicast_send(input[1], message)
		}
	}
}