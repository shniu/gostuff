package options

import (
	"github.com/stretchr/testify/assert"
	"path"
	"strings"
	"testing"
)

func TestKvsOptions_GetDataFilePath(t *testing.T) {
	filePath := Options.GetDataFilePath(1)
	t.Log(filePath)
	assert.True(t, strings.LastIndex(filePath, Options.DataFileSuffix) > 0)
}

func TestKvsOptions_GetFpFunc(t *testing.T) {
	fPath := Options.GetFpFunc(true)(0)
	t.Log(fPath)
	assert.True(t, strings.LastIndex(fPath, Options.DataFileSuffix) > 0)

	fPath = Options.GetFpFunc(false)(0)
	t.Log(fPath)
	assert.True(t, strings.LastIndex(fPath, "000.hint") > 0)
}

func TestKvsOptions_GetLockFilePath(t *testing.T) {
	p := path.Join(Options.DataDir, strings.Join([]string{"123", Options.DataFileSuffix}, ""))
	t.Log(Options.DataDir)
	t.Log(Options.DataFileSuffix)
	t.Log(p)

	maxSize := 1 << 26
	t.Log(maxSize, "B")
	t.Log(maxSize/1024, "KB")           // KB = 1024 Byte
	t.Log(maxSize/1024/1024, "MB")      // MB = 1024 KB
	t.Log(maxSize/1024/1024/1024, "GB") // GB = 1024 MB
}
