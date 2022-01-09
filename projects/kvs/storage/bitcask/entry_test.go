package bitcask

import (
	"log"
	"testing"
)

func TestNewEntry(t *testing.T) {
	var e = entry{}
	log.Println(e)
	//log.Println(e == nil)

	b := make([]byte, 5)
	copy(b[:5], []byte("hello"))
	copy(b[5:5], []byte("world"))
	log.Println(string(b))

	const (
		a = 1 << iota
	)
	log.Println(a)
	bb := uint16(0)
	bb |= a
	log.Println(bb)
}
