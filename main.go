package main

import (
	"fmt"
	"log"

	"github.com/codecrafters-io/kafka-starter-go/app/server"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// address
	addr := "0.0.0.0:9092"

	if err := server.Listen(addr); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}
}
