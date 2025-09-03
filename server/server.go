package server

import (
	"encoding/binary"
	"fmt"
	"net"
)

func Listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to bind to %s: %w", addr, err)
	}
	defer ln.Close()

	fmt.Println("Server listening on", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)
	if _, err := conn.Read(buf); err != nil {
		fmt.Println("read error: ", err)
		return
	}

	//Proceeding Kafka request
	apiKey := int16(binary.BigEndian.Uint16(buf[4:6]))
	version := int16(binary.BigEndian.Uint16(buf[6:8]))
	corrId := int32(binary.BigEndian.Uint32(buf[8:12]))

	// Calling api key handler
	resp, err := HandleRequest(apiKey, corrId, version)
	if err != nil {
		fmt.Println("handler error: ", err)
		return
	}

	if _, err := conn.Write(resp); err != nil {
		fmt.Println("write error: ", err)
	}

}
