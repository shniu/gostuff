package cache

const KeyMaxLength = 1024

// Cache interface
type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() []byte
}
