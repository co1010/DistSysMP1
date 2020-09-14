package main

import (
	"bufio"
	"os"
	"strings"
	"time"
)

func main() {
	startProcesses()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		input := strings.Split(text, " ")
		content := strings.Join(input[2:], " ")
		content = strings.TrimRight(content, "\r\n")
		if input[0] == "send" {
			message := Message{content, time.Now()}
			go unicast_send(input[1], message)
		}
	}
}