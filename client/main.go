package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net"
)

func main() {
	client, err := net.Dial("tcp", ":5001")
	defer client.Close()
	for {
		if err != nil {
			log.Fatalf("COuld not start client")
		}
		number := rand.Int()
		client.Write([]byte(fmt.Sprintf("pong: %d", number)))
		data := make([]byte, 1024)
		n, err := client.Read(data)
		if err != nil {
			log.Fatalf("could not read from server")
		}
		fmt.Printf("Read %d bytes: %s\n", n, data)
	}
}
