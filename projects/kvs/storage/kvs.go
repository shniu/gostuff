package storage

import (
	"github.com/shniu/gostuff/projects/kvs/errors"
	"github.com/shniu/gostuff/projects/kvs/options"
	"github.com/shniu/gostuff/projects/kvs/storage/bitcask"
)

// Closer is the interface that wraps the basic Close method.
type Closer interface {
	Close() error
}

// Iterator, only the leveldb storage engine support for now
type Iterator interface {
	// Todo
}

// Kvs is the interface to the key-value storage system
type Kvs interface {
	Get(key string) []byte

	// Same as `Set`, WriteOptions will be added later
	Put(key string, value []byte) (bool, error)
	Delete(key string) (bool, error)

	Closer
	Iterator

	// Merge
	Merge() (bool, error)
}

// Open database
func Open(opts *options.KvsOptions) (Kvs, error) {
	var kvsImpl Kvs
	var err error

	switch opts.EngineType {
	case options.Bitcask:
		kvsImpl, err = bitcask.Open(opts)
	case options.Leveldb:
		kvsImpl, err = bitcask.Open(opts)
	default:
		err = errors.EngineTypeNotFound
	}

	if err != nil {
		go merge()
	}

	return kvsImpl, err
}

func merge() {
    // TODO merge function
}
