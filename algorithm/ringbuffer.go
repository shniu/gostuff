package algorithm

import (
	"github.com/shniu/gostuff/kvs/errors"
)

// 实现一个 RingBuffer
// fill count
type RingBuffer struct {
	// 写指针
	writePos int
	// 可读元素个数
	available int
	// 容量
	capacity int

	// 数组
	elements []int
}

func NewRingBuffer(capacity int) (*RingBuffer, error) {
	if capacity <= 0 {
		return nil, errors.New("Capacity must be greater than zero.")
	}

	return &RingBuffer{
		writePos: 0,
		capacity: capacity,
		elements: make([]int, capacity),
	}, nil
}

func (r *RingBuffer) Put(element int) bool {
	if r.notFull() {
		if r.writePos >= r.capacity {
			r.writePos = 0
		}

		r.elements[r.writePos] = element
		r.writePos++
		r.available++
		return true
	}

	return false
}
func (r *RingBuffer) notFull() bool {
	return r.available < r.capacity
}

func (r *RingBuffer) Take() (int, bool) {
	if r.isEmpty() {
		return -1, false
	}

	// readPos
	readPos := r.writePos - r.available
	if readPos < 0 {
		readPos += r.capacity
	}
	var ele = r.elements[readPos]
	r.available--
	return ele, true
}

func (r *RingBuffer) isEmpty() bool {
	return r.available == 0
}
