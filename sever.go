package main

import "fmt"

func unicast_receive(source string, message Message){
	fmt.Printf("Received \"%s\" from process %s\n", message.Content, source)
}