package bitcask

import "sync"

var rwLock *sync.RWMutex
var keyDirsOnce sync.Once

var keyDirs *KeyDirs

func init() {
	rwLock = &sync.RWMutex{}
}

type KeyDirs struct {
	// Hash map, an in-memory structure, to retrieve the key
	entries map[string]*entry
}

func NewKeyDirs() *KeyDirs {
	rwLock.Lock()
	defer rwLock.Unlock()

	// todo ?
	keyDirsOnce.Do(func() {
		if keyDirs == nil {
			keyDirs = &KeyDirs{
				entries: make(map[string]*entry),
			}
		}
	})

	return keyDirs
}

func (kd *KeyDirs) get(key string) *entry {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return kd.entries[key]
}

func (kd *KeyDirs) put(key string, e *entry) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	kd.entries[key] = e
}

func (kd *KeyDirs) del(key string) {
	rwLock.Lock()
	defer rwLock.Unlock()
	delete(kd.entries, key)
}

func (kd *KeyDirs) Keys() []string {
	keys := make([]string, 0, len(kd.entries))
	for k, v := range kd.entries {
		if v.flag&RecordDeleteFlag == 0 {
			keys = append(keys, k)
		}
	}

	return keys
}
