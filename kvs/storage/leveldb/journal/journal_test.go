package journal

import (
	"bytes"
	"testing"
)

func TestNewReader(t *testing.T) {
	NewReader(nil, nil, true, true)
}

func TestNewWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	w.Next()
}
