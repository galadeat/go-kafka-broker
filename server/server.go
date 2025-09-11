package server

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

func ListenandServe(addr string) error {
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(10 * time.Second))
	for {
		buf, err := createBuffer(conn)
		if err != nil {
			fmt.Println(" incorrect request")
			return
		}

		//Proceeding Kafka request
		if len(buf) < 12 {
			return
		}
		apiKey := int16(binary.BigEndian.Uint16(buf[:2]))
		version := int16(binary.BigEndian.Uint16(buf[2:4]))
		corrId := int32(binary.BigEndian.Uint32(buf[4:8]))

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

}

func createBuffer(conn net.Conn) ([]byte, error) {
	msgLength := make([]byte, 4)
	if _, err := io.ReadFull(conn, msgLength); err != nil {
		fmt.Printf("error: can't read message size %s", err)
		return nil, err
	}
	msgLen := binary.BigEndian.Uint32(msgLength)
	if msgLen > MaxLenRequest {
		return nil, fmt.Errorf("request is too long")
	}

	buffLength := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, buffLength); err != nil {
		fmt.Printf("error: can't read message size %s", err)
		return nil, err
	}
	return buffLength, nil
}
