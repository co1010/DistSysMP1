package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
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

func startProcesses() {
	filelines := readFile("config.txt")
	for _, line := range filelines[1:] {
		process := strings.Split(line, " ")
		ch := make(chan net.Conn)
		go startServer(process[0], process[1], ":"+process[2], ch)
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

func startServer(id, ip, port string, ch chan net.Conn) {
	ln, err := net.Listen("tcp", port)
	checkError(err)
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		checkError(err)
		decoder := gob.NewDecoder(conn)
		var message Message
		decoder.Decode(&message)
		unicast_receive(conn.RemoteAddr().String(), message)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	netData, err := bufio.NewReader(c).ReadString('\n')
	checkError(err)

	temp := strings.TrimSpace(netData)
	fmt.Println(temp)
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
