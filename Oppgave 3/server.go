package main

import (
	"fmt"
	"net"
	"bufio"
	//"strings"
	"./encryption"
)// only needed below for sample processing

func main() {

	// Server

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		encryptedMessage, _ := bufio.NewReader(conn).ReadString('\n')

		convert := []byte (encryptedMessage)

		test := encryption.DecryptMessage(convert)

		fmt.Printf("Message received: %s", test)

	}
}