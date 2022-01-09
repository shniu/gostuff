package cache

func NewCache(cType string) Cache {

	if cType == "inMemoryCache" {
		return newInMemoryCache()
	}

	panic("unknown cache type " + cType)
}
