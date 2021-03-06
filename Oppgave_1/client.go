/**
Hentet fra: https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/
 */
package main

import (
	"fmt"
	"net"
	"time"
)

/**
Lager en client som sender meldingen "Møte Fr 5.5 14:45 Flåklypa" til "server.go"
 */
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
		buf := []byte(msg)
		_,err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}

func checkError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}