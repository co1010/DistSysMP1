package main

import (
	"MP1/utils"
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

// Sends a message to specified destination ID
func unicast_send(destination string, message utils.Message) {

	// Use util function to grab IP and PORT of the destination
	ip, port, err := utils.GetNodeDetails(destination)
	utils.CheckError(err)
	CONNECT := ip+":"+port

	// Connect to IP:PORT of destination
	c, err := net.Dial("tcp", CONNECT)
	utils.CheckError(err)

	// Encode and send message
	encoder := gob.NewEncoder(c)
	encoder.Encode(message)

	// Print that the message has been sent
	fmt.Printf("Sent \"%s\" to process %s at %s\n", message.Content, destination, time.Now())
}
