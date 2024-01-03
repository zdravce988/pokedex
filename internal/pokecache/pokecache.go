package pokecache

import (
    "time"
    "sync"
)

type Cache struct {
    entries map[string]cacheEntry
    interval time.Duration
    mu *sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

func NewCache(interval time.Duration) Cache {
    cache := Cache{
        entries: make(map[string]cacheEntry),
        interval: interval,
        mu: &sync.Mutex{},
    }
    cache.entries = make(map[string]cacheEntry)
    go cache.reapLoop()

    return cache
}

func (c Cache) Add(key string, val []byte) {
    newEntry := cacheEntry{createdAt: time.Now(), val: val}
    c.entries[key] = newEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
    value, exists := c.entries[key]
    if !exists {
        return nil, false
    }
    return value.val, true
}

func (c Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    for {
        select {
        case <- ticker.C:
            c.mu.Lock()
            for key, entry:= range c.entries {
                if time.Now().Sub(entry.createdAt) > c.interval {
                    delete(c.entries, key)
                }
            }
            c.mu.Unlock()
        }
    }
}
