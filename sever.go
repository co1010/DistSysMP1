package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func unicast_receive(source *string, message Message) {
	fmt.Printf("Received \"%s\" from process %s at %s\n", message.Content, *source, time.Now())
}

func startServer(node *string) {
	filelines := readFile("config.txt")
	port := ""
	for _, line := range filelines[1:] {
		process := strings.Split(line, " ")
		if process[0] == *node {
			port = ":"+process[2]
		}
	}
	ln, err := net.Listen("tcp", port)
	checkError(err)
	defer ln.Close()
	conn, err := ln.Accept()
	checkError(err)

	min, max := getDelayParams()
	delay := rand.Intn(max-min)+min
	time.Sleep(time.Duration(delay) * time.Millisecond)

	decoder := gob.NewDecoder(conn)
	var message Message
	decoder.Decode(&message)
	unicast_receive(node, message)
}