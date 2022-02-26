package protocol

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	header := Header([11]byte{})
	header[0] = MagicNumber()

	fmt.Println(header)

	fmt.Println(MessageType(0x01))
}
