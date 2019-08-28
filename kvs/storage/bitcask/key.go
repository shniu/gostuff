package bitcask

import (
	"unicode"
)

const (
	MaxKeyLen = 250
)

// Key valid
func IsValidKey(key string) bool {
	length := len(key)
	if length == 0 || length > MaxKeyLen {
		logger.Printf("bad key, length=%d\n", length)
		return false
	}

	if key[0] <= ' ' || key[0] == '@' || key[0] == '?' {
		logger.Printf("bad key, length=%d, key[0]=%c\n", length, key[0])
		return false
	}

	for _, r := range key {
		if unicode.IsControl(r) || unicode.IsSpace(r) {
			logger.Printf("bad key, length=%d, %s\n", length, key)
			return false
		}
	}

	return true
}
