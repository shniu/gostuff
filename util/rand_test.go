package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	randStr := RandStringBytesMaskImprSrc(1024)

	t.Log(randStr)
	assert.Equal(t, 1024, len(randStr))
}
