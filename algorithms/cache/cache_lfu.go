package cache

import "container/list"

type CacheItem struct {
	key             string
	value           interface{}
	frequencyParent *list.Element
}

type FrequencyItem struct {
	entries map[*CacheItem]byte
	freq    int
}

type Cache struct {
	bykey    map[string]*CacheItem
	freqs    *list.List
	capacity int
	size     int
}

func New() *Cache {
	cache := new(Cache)
	cache.bykey = make(map[string]*CacheItem)
	cache.freqs = list.New()
	cache.size = 0
	cache.capacity = 100

	return cache
}

func (cache *Cache) Set(key string, value interface{}) {
	if item, ok := cache.bykey[key]; ok {
		item.value = value
	} else {
		item := new(CacheItem)
		item.key = key
		item.value = value
		cache.bykey[key] = item
		cache.size++
	}
}

func (cache *Cache) Get(key string) interface{} {
	if e, ok := cache.bykey[key]; ok {
		return e.value
	}
	return nil
}
