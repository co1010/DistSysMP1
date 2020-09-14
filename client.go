package main

import (
	"encoding/gob"
	"net"
)

func unicast_send(destination string, message Message) {
	ip, port := getNodeDetails(destination)
	CONNECT := ip+":"+port
	c, err := net.Dial("tcp", CONNECT)
	checkError(err)
	encoder := gob.NewEncoder(c)
	encoder.Encode(message)
}
