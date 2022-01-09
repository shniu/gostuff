package util

import "hash/crc32"

// https://golang.org/pkg/hash/crc32/
// crc wiki: https://en.wikipedia.org/wiki/Cyclic_redundancy_check

// Get crc
func Crc(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
