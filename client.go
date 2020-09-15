package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func unicast_send(destination string, message Message) {
	fmt.Printf("Sent \"%s\" to process %s at %s\n", message.Content, destination, time.Now())
	ip, port := getNodeDetails(destination)
	CONNECT := ip+":"+port
	c, err := net.Dial("tcp", CONNECT)
	checkError(err)
	encoder := gob.NewEncoder(c)
	encoder.Encode(message)
}
