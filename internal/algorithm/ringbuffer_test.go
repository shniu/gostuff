package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestRingBuffer(t *testing.T) {
	rb, _ := NewRingBuffer(5)
	rb.Put(1)
	rb.Put(30)
	rb.Put(2)
	rb.Put(33)
	rb.Put(40)

	v, ok := rb.Take()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, v)
	fmt.Println(rb)

	v, ok = rb.Take()
	assert.Equal(t, true, ok)
	assert.Equal(t, 30, v)
	fmt.Println(rb)

	v, ok = rb.Take()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, v)
	fmt.Println(rb)

	rb.Put(100)
	rb.Take()
	rb.Take()
	v, ok = rb.Take()
	assert.Equal(t, true, ok)
	assert.Equal(t, 100, v)
	fmt.Println(rb)

	v, ok = rb.Take()
	assert.Equal(t, false, ok)
	assert.Equal(t, -1, v)
	fmt.Println(rb)
}

func TestRingBuffer_init(t *testing.T) {
	_, err := NewRingBuffer(0)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Capacity must be greater than zero.")
}
