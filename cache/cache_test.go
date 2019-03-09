package cache

import (
	"github.com/shniu/cache/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func Test_InMemoryCache_Set_Succeed(t *testing.T) {
	cache := &inMemoryCache{
		entries: make(map[string][]byte),
	}
	err := cache.Set("uid:889192", []byte("{\"username\": \"Little Q\"}"))

	assert.Nil(t, err)
	assert.Equal(t, 1, len(cache.entries))
	assert.Equal(t, []byte("{\"username\": \"Little Q\"}"), cache.entries["uid:889192"])
}

func TestInMemoryCache_Set_Given_emptyKey_Then_error(t *testing.T) {
	cache := &inMemoryCache{
		entries: make(map[string][]byte),
	}
	err := cache.Set("", []byte(""))

	assert.NotNil(t, err)
	assert.Equal(t, "empty key is not allowed", err.Error())
	assert.Equal(t, 0, len(cache.entries))
}

func TestInMemoryCache_Set_Given_GtKeyMaxSize_Then_error(t *testing.T) {
	key1024 := util.RandStringBytesMaskImprSrc(1024)
	cache := &inMemoryCache{
		entries: make(map[string][]byte),
	}
	err := cache.Set(key1024, []byte(""))

	assert.Nil(t, err)

	keyGt1024 := util.RandStringBytesMaskImprSrc(1025)
	err2 := cache.Set(keyGt1024, []byte(""))
	assert.NotNil(t, err2)
}

func TestInMemoryCache_Get_Exist_key_Then_succeed(t *testing.T) {
	cache := &inMemoryCache{
		entries: make(map[string][]byte),
	}

	key := "uuu"
	expectedVal := []byte("hello world")
	_ = cache.Set(key, expectedVal)
	value, err := cache.Get(key)

	assert.Nil(t, err)
	assert.Equal(t, expectedVal, value)
}

func TestInMemoryCache_Get_NotExistKey_Then_error(t *testing.T) {
	cache := &inMemoryCache{
		entries: make(map[string][]byte),
	}

	value, err := cache.Get("not_exists")

	assert.NotNil(t, err)
	assert.Equal(t, "the key does not exist", err.Error())
	assert.Nil(t, value)
}
