/**
Hentet fra: https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/
 */
package main

import (
	"fmt"
	"net"
	"time"
	"./encryption"
)

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
	checkError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	checkError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	checkError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := "Møte Fr 5.5 14:45 Flåklypa"
		i++
		encryptedMessage := encryption.EncryptMessage(msg)

		//buf := []byte(msg)
		_,err := Conn.Write(encryptedMessage)
		if err != nil {
			fmt.Println(encryptedMessage, err)
		}
		time.Sleep(time.Second * 1)
	}
}

func checkError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}