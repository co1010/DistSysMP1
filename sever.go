package main

import (
	"MP1/utils"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// Receive the message and print it
func unicast_receive(source *string, message utils.Message) {
	fmt.Printf("Received \"%s\" from process %s at %s\n", message.Content, *source, time.Now())
}

// Initialize TCP listen server
func startServer(node *string) {

	// Set port using util function
	_, port, err := utils.GetNodeDetails(*node)
	utils.CheckError(err)
	ln, err := net.Listen("tcp", ":"+port)
	utils.CheckError(err)
	defer ln.Close()
	conn, err := ln.Accept()
	utils.CheckError(err)

	// Use goroutine to implement delay to avoid bottleneck
	go decodeMessage(node, conn)
}

// Delay message to simulate network delay and then decode it
func decodeMessage(node *string, conn net.Conn) {

	// Get delay params using util function
	min, max := utils.GetDelayParams()
	delay := rand.Intn(max-min)+min
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// Decode conn and call unicast_receive
	decoder := gob.NewDecoder(conn)
	var message utils.Message
	decoder.Decode(&message)
	unicast_receive(node, message)
}