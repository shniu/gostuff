package tests

// Just for test diskv (https://github.com/peterbourgon/diskv)
// diskv: is a simple, persistent key-value store written in the Go
//   language. It starts with an incredibly simple API for storing arbitrary data on
//   a filesystem by key, and builds several layers of performance-enhancing
//   abstraction on top.  The end result is a conceptually simple, but highly
//   performant, disk-backed storage system.

import (
	"crypto/md5"
	"fmt"
	"github.com/peterbourgon/diskv"
	"io"
	"testing"
)

func TestDiskv(t *testing.T) {
	t.Log("Test diskv")

	transformFunc := func(s string) []string {
		return []string{}
	}

	d := diskv.New(diskv.Options{
		BasePath:     "/tmp/diskv",
		CacheSizeMax: 1024 * 1024,
		Transform:    transformFunc,
	})

	key := "alpha"
	d.Write(key, []byte{'1', '2', '3'})
	d.Write(key, []byte{'a', 'b', 'c'})

	value, _ := d.Read(key)
	t.Logf("key=%s, value=%v\n", key, value)

}

func TestDiskvAdvance(t *testing.T) {
	t.Log("Test diskv advanced feature")

	// TransformFunc
	transformFunc := func(s string) *diskv.PathKey {
		h := md5.New()
		io.WriteString(h, s)

		md5Str := fmt.Sprintf("%x", h.Sum(nil))
		//t.Log(md5Str)

		return &diskv.PathKey{
			Path:     []string{},
			FileName: md5Str,
		}
	}

	inverseTransformFunc := func(pk *diskv.PathKey) (key string) {
		return pk.FileName
	}

	d := diskv.New(diskv.Options{
		BasePath:          "/tmp/diskv",
		AdvancedTransform: transformFunc,
		InverseTransform:  inverseTransformFunc,
		CacheSizeMax:      1024 * 1024,
	})

	key := "aaa"
	d.Write(key, []byte{'e', 'e', 'e'})
	value, _ := d.Read(key)
	t.Logf("key=%s, value=%v\n", key, value)
}
