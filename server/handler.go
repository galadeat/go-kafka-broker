package server

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/app/handlers"
	"github.com/codecrafters-io/kafka-starter-go/app/protocol"
)

// API handler signature
type HandlerFunc func(correlationID int32, version int16) ([]byte, error)

// registered handlers
var Registry = map[int16]HandlerFunc{
	protocol.APIKeyApiVersions: handlers.ApiVersionsHandler,
}

func HandleRequest(apiKey int16, correlationID int32, version int16) ([]byte, error) {
	handler, ok := Registry[apiKey]
	if ok {
		return handler(correlationID, version)
	}
	fmt.Printf("Unknown API key: %d\n", apiKey)

	msg := protocol.Message{
		Header: protocol.Header{
			CorrelationID: correlationID,
			ErrorCode:     protocol.ErrorUnsupportedVersion,
		},
		ArrayLength: 0,
		Body:        nil,
	}
	return msg.Encode()
}
