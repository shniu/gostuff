package main

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestKvsdb_Put(t *testing.T) {
	kvsdb := Open("/tmp/kvs", nil)
	err := kvsdb.Put([]byte("name"), []byte("leveldb"))
	t.Log(err)
}

func TestKvsdb_Get(t *testing.T) {
	db, err := leveldb.OpenFile("/tmp/leveldb", nil)
	if err != nil {
		t.Error("Open /tmp/leveldb failed.")
	}
	err = db.Put([]byte("k1"), []byte("v1"), nil)
	value, err := db.Get([]byte("k1"), nil)
	t.Logf("k1=%s", value)

	err = db.Close()
	if err != nil {
		t.Error("Close db failed.")
	}
}
