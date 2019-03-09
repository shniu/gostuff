package cache

import "testing"

func Test_InMemory(t *testing.T) {
	cache := &InMemoryImpl{}
	cache.Set([]byte("abc"), []byte("{\"username\": \"Little Q\"}"))
}
