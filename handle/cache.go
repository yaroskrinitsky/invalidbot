package handle

import "time"

//CacheEntry ...
type CacheEntry struct {
	val       string
	timestamp time.Time
}

//Cache ...
type Cache struct {
	col map[string]CacheEntry
	//objects lifetime in minutes
	lifetime time.Duration
}

//Get ...
func (c *Cache) Get(key string) (string, bool) {
	entry := c.col[key]

	if entry.val != "" && time.Now().Sub(entry.timestamp) <= c.lifetime*time.Minute {
		return entry.val, true
	}

	return "", false
}

//Add ...
func (c *Cache) Add(key string, val string) {
	ts := time.Now()
	c.col[key] = CacheEntry{
		val:       val,
		timestamp: ts,
	}
}
