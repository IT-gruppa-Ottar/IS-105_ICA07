package main

import (
	"fmt"
	"net"
	"os"
	"./encryption"
)



func main() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	n,addr,err := ServerConn.ReadFromUDP(buf)

	decryptedMessage := encryption.DecryptMessage(buf)

	fmt.Println("Received ",string(decryptedMessage[0:n]), " from ",addr)

	if err != nil {
		fmt.Println("Error: ",err)
	}
}

/* A Simple function to verify error */
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}