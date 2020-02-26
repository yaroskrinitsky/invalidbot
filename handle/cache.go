package handle

import "time"

//CacheEntry represents a unit of data in cache
type CacheEntry struct {
	val       string
	timestamp time.Time
}

//Cache stores execution results for the command service
type Cache struct {
	col map[string]CacheEntry
	//objects lifetime in minutes
	lifetime time.Duration
}

//NewCache creates a new cache object
func NewCache(duration int) Cache {
	return Cache{
		col:      make(map[string]CacheEntry),
		lifetime: time.Duration(duration) * time.Minute,
	}
}

//Get retrieves a cache entry if it exists, otherwise it returns an empty string and false
func (c *Cache) Get(key string) (string, bool) {
	entry := c.col[key]

	if entry.val != "" && time.Now().Sub(entry.timestamp) <= c.lifetime*time.Minute {
		return entry.val, true
	}

	return "", false
}

//Add puts a new value to the cache by the provided key
func (c *Cache) Add(key string, val string) {
	ts := time.Now()
	c.col[key] = CacheEntry{
		val:       val,
		timestamp: ts,
	}
}
