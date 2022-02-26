package protocol

import "encoding/binary"

// This is a custom message protocol.
// BitOffset        0-15        ｜     16-47    ｜     46-63      ｜       64-71      ｜      72-79      ｜   80-87
//  0-8           Magic Number     Total Length    Header Length     Protocol Version    Message Type      Serializer Type
//  88-103        Message ID
//  Extend header Extend fields of header
//  Content       Payload

// Magic number
const magicNumber byte = 0xff

func MagicNumber() byte {
	return magicNumber
}

type TotalLength int32
type HeaderLength int16
type Version byte

type MessageType byte
const (
	// Request Type
	Request MessageType = iota
	// Response Type
	Response
)

type SerializerType byte
const (
	// Using raw []byte and don't serialize/deserialize
	SerializerNone SerializerType = iota
	// JSON for Payload
	JSON
	// Protobuf for Payload
	ProtoBuffer
	// MsgPack for Payload
	MsgPack
	// Thrift for Payload
	Thrift
)

// Extend header
//  compressor
type CompressType byte

const (
	CompressNone CompressType = iota
	Gzip
)

// Define message
type Message struct {

}

type Header [11]byte

func (h Header) CheckMagicNumber() bool {
	return h[0] == magicNumber
}

func (h Header) Version() byte {
	return h[1]
}

func (h Header) TotalLength() TotalLength {
	return TotalLength(binary.BigEndian.Uint32(h[2:6]))
}
