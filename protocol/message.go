package protocol

import (
	"encoding/binary"
)

type ApiVersion struct {
	ApiKey       int16
	MinSupported int16
	MaxSupported int16
	TagBuffer    int8
}

type Header struct {
	CorrelationID int32
	ErrorCode     int16
}

type Message struct {
	MessageSize  int32
	Header       Header
	ArrayLength  int8
	Body         []ApiVersion
	ThrottleTime int32
	TagBuffer    int8
}

func (m *Message) Encode() ([]byte, error) {

	body, err := m.encodeBody()
	if err != nil {
		return nil, err
	}

	var msg []byte
	msg, _ = binary.Append(msg, binary.BigEndian, int32(len(body)))
	msg = append(msg, body...)
	//For debug purpose
	//fmt.Println(msg)
	return msg, nil

}

func (m *Message) encodeBody() (buf []byte, err error) {
	buf, err = binary.Append(buf, binary.BigEndian, m.Header.CorrelationID)
	if err != nil {
		return nil, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, m.Header.ErrorCode)
	if err != nil {
		return nil, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, m.ArrayLength)
	if err != nil {
		return nil, err
	}

	for _, b := range m.Body {
		buf = append(buf, b.serializeStruct()...)

	}

	buf, err = binary.Append(buf, binary.BigEndian, m.ThrottleTime)
	if err != nil {
		return nil, err
	}

	buf, err = binary.Append(buf, binary.BigEndian, m.TagBuffer)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (apiStruct *ApiVersion) serializeStruct() (binarr []byte) {
	binarr, _ = binary.Append(binarr, binary.BigEndian, apiStruct.ApiKey)
	binarr, _ = binary.Append(binarr, binary.BigEndian, apiStruct.MinSupported)
	binarr, _ = binary.Append(binarr, binary.BigEndian, apiStruct.MaxSupported)
	binarr, err := binary.Append(binarr, binary.BigEndian, apiStruct.TagBuffer)
	if err != nil {
		return nil
	}
	return binarr
}
