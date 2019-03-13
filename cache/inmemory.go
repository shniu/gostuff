package cache

import (
	"errors"
	"sync"
)

type inMemoryCache struct {
	// https://blog.golang.org/go-maps-in-action
	entries map[string][]byte

	mutex sync.RWMutex
	Stat
}

func newInMemoryCache() Cache {
	return &inMemoryCache{
		entries: make(map[string][]byte),
	}
}

func (mc *inMemoryCache) Set(key string, value []byte) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	if key == "" {
		return errors.New("empty key is not allowed")
	}

	if len(key) > KeyMaxLength {
		return errors.New("the key is greater than 1024")
	}

	mc.entries[key] = value
	return nil
}

func (mc *inMemoryCache) Get(key string) ([]byte, error) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	if key == "" {
		return nil, errors.New("empty key is not allowed")
	}

	if value, ok := mc.entries[key]; ok {
		return value, nil
	}

	return nil, errors.New("the key does not exist")
}

func (mc *inMemoryCache) Del(key string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	delete(mc.entries, key)
	return nil
}

func (mc *inMemoryCache) GetStat() Stat {
	return mc.Stat
}
