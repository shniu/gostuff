package util

import "bytes"

// Combine multi byte slice
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
