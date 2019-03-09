package cache

// Cache interface
type Cache interface {
	Set(key []byte, value []byte) error
	Get(key []byte) []byte
	Del(key []byte) error
	GetState() []byte
}

type InMemoryImpl struct {
}

func (memory *InMemoryImpl) Set(key []byte, value []byte) error {
	return nil
}
