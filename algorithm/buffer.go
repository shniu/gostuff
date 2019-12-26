package algorithm

import (
	"sync"
	"github.com/shniu/gostuff/kvs/errors"
	"fmt"
)

// 更加通用的 RingBuffer 实现，并支持以下接口：
// - io.Reader
// - io.Writer
// - io.ByteReader
// - io.ByteWriter
type Ring struct {
	buf       []byte
	size      int
	wp        int // write pos
	available int // available to read

	m sync.Mutex // lock
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (r *Ring) Read(p []byte) (n int, err error) {
	lengthToRead := len(p)
	if lengthToRead == 0 {
		return 0, nil
	}

	r.m.Lock()
	if r.isEmpty() {
		r.m.Unlock()
		return 0, errors.New("No data to read.")
	}

	rp := r.readPos()
	bytesCanRead := min(r.available, lengthToRead)
	if rp+bytesCanRead < r.size {
		copy(p, r.buf[rp:rp+bytesCanRead])
	} else {
		copy(p, r.buf[rp:])
		remain := rp + bytesCanRead - r.size
		copy(p[r.size-rp:], r.buf[0:remain])
	}
	r.available -= bytesCanRead
	r.m.Unlock()
	return bytesCanRead, nil
}

func (r *Ring) isEmpty() bool {
	return r.available == 0
}

func (r *Ring) ReadByte() (byte, error) {
	r.m.Lock()
	if r.available == 0 {
		r.m.Unlock()
		return 0, errors.New("No data to read.")
	}

	rp := r.readPos()
	b := r.buf[rp]
	r.m.Unlock()
	return b, nil
}

func (r *Ring) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	r.m.Lock()
	if r.isFull() {
		r.m.Unlock()
		return 0, errors.New("No space to write.")
	}

	bytesCanWrite := min(r.size - r.available, len(p))
	fmt.Println(bytesCanWrite)

	r.m.Unlock()
	return 0, nil
}
func (r *Ring) isFull() bool {
	return r.available == r.size
}

func (r *Ring) WriteByte(c byte) error {
	return nil
}
func (r *Ring) readPos() int {
	rp := r.wp - r.available
	if rp < 0 {
		rp += r.size
	}
	return rp
}

func NewRing(size int) *Ring {
	return &Ring{
		buf:  make([]byte, size),
		size: size,
	}
}
