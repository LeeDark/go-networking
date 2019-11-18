package main

import (
	"fmt"
	"net"
)

func main() {
	service := ":1201"

	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleIPClientThreaded(conn)
	}
}

func handleIPClientThreaded(conn net.Conn) {
	// close connection on exit
	defer conn.Close()

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		fmt.Println(string(buf[0:n]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}