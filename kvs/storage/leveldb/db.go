package main

import (
	"github.com/shniu/gostuff/kvs/options"
	"github.com/shniu/gostuff/kvs/storage/leveldb/journal"
	"github.com/shniu/gostuff/kvs/storage/leveldb/memdb"
)

// DB
type Kvsdb struct {

	// MemDB
	journal *journal.Writer
	mem     memdb.DB
}

func Open(path string, o *options.KvsOptions) *Kvsdb {
	// Todo leveldb/options?
	return nil
}

// Get
func (db *Kvsdb) Get(key []byte) (value []byte, err error) {
	// Todo
	return nil, nil
}

// Has returns true if the DB does contains the given key
func (db *Kvsdb) Has(key []byte) bool {
	return true
}

// Put
func (db *Kvsdb) Put(key []byte, value []byte) error {
	// Todo
	// Write 操作
	//   日志数据结构 journal
	//   memTable 数据结构
	return nil
}

// Delete
func (db *Kvsdb) Delete(key []byte) error {
	// Todo
	return nil
}

// Close
func (db *Kvsdb) Close() {

}
