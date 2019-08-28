package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

func TestKvsdb_Put(t *testing.T) {
	kvsdb := Open("/tmp/kvs", nil)
	err := kvsdb.Put([]byte("name"), []byte("leveldb"))
	t.Log(err)
}

func TestKvsdb_Get(t *testing.T) {
	db, err := leveldb.OpenFile("/tmp/leveldb", nil)
	if err != nil {
		logger.Error("Open /tmp/leveldb failed.")
	}
	err = db.Put([]byte("k1"), []byte("v1"), nil)
	value, err := db.Get([]byte("k1"), nil)
	logger.Infof("k1=%s", value)

	err = db.Close()
	if err != nil {
		logger.Error("Close db failed.")
	}
}
