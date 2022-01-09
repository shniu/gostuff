package storage

import (
	"os"
	"testing"
)

// For test the file in golang
// 磁盘 IO：https://tech.meituan.com/2017/05/19/about-desk-io.html

// Working with Files in Go： @link https://www.devdungeon.com/content/working-files-go

func Test_File(t *testing.T) {
	// Golang 中最基础的File抽象 os.File
	var newFile *os.File
	var err error

	newFile, err = os.Create("test.txt")
	if err != nil {
		t.Log("Create test.txt error")
	}
	t.Log(newFile)

	fileInfo, err := newFile.Stat()
	fileInfo.Name()
	fileInfo.IsDir()
	newFile.Close()
}

// File 相关的包
// ioutil
// io
// bufio
// The bufio package lets you create a buffered writer so you can work with a buffer in memory before
// writing it to disk. This is useful if you need to do a lot manipulation on the data before writing
// it to disk to save time from disk IO. It is also useful if you only write one byte at a time and want
// to store a large number in memory before dumping it to file at once, otherwise you would be performing
// disk IO for every byte. That puts wear and tear on your disk as well as slows down the process.

// zip
// filepath
// gzip
// http get file
// crypto hash

// @link golang pkg: https://golang.org/pkg/
// @link golang doc: https://golang.org/doc/
