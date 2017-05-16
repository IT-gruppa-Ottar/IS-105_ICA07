/**
Hentet fra: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
*/
package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"./encryption"
)

func main() {
	// Client
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		text, _ := reader.ReadString('\n')

		convert := encryption.EncryptMessage(text)
		send := string(convert)

		// send to socket
		fmt.Fprintf(conn, send + "\n")
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message from server: "+message)
	}
}
