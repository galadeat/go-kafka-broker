package handlers

import "github.com/codecrafters-io/kafka-starter-go/app/protocol"

// Api VersionsHandler формирует ответ на ApiVersions запрос
func ApiVersionsHandler(correlationID int32, version int16) ([]byte, error) {

	// List of supported APIs
	supported := []protocol.ApiVersion{
		{ApiKey: protocol.APIKeyApiVersions, MinSupported: 0, MaxSupported: 4, TagBuffer: 0},
	}

	msg := protocol.Message{
		Header:       protocol.Header{CorrelationID: correlationID},
		ArrayLength:  int8(len(supported) + 1), // Compact Array Length
		Body:         supported,
		ThrottleTime: 0,
		TagBuffer:    0,
	}

	if !protocol.IsSuportedVersion(version) {
		msg.Header.ErrorCode = protocol.ErrorUnsupportedVersion
	}

	return msg.Encode()
}
