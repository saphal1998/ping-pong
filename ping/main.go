package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 1024)
	for {
		_, err := conn.Read(data)
		if err != nil {
			fmt.Printf("Closed connection")
			return
		}
		fmt.Println(fmt.Sprintf("%s", data))
		conn.Write([]byte(fmt.Sprintf("server sends back %s", data)))
	}

}

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("could not accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
