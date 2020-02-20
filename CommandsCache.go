package main

import "time"

type CacheEntry struct {
	val string
	timestamp time.Time
}

type CommandsCache struct {
	col map[string]CacheEntry
	//objects lifetime in minutes
	lifetime time.Duration
}

func(cc *CommandsCache) Get(key string) (string, bool) {
	entry := cc.col[key]

	if entry.val != "" && time.Now().Sub(entry.timestamp) <= cc.lifetime*time.Minute {
		return entry.val, true
	}

	return "", false
}

func(cc *CommandsCache) Add(key string, val string) {
	ts := time.Now()
	cc.col[key] = CacheEntry{
		val:       val,
		timestamp: ts,
	}
}