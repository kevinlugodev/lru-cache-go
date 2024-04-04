package lrucache

import (
	"container/list"
	"time"
)

type LRUCache struct {
	capacity *uint
	data     map[string]*list.Element
	lru      *list.List
}

type entry struct {
	key            string
	value          interface{}
	expirationTime *time.Time
}

// NewLRUCache returns a new LRUCache with the given capacity.
//
// It takes a pointer to uint as a parameter and returns a pointer to LRUCache.
func NewLRUCache(capacity *uint) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		data:     make(map[string]*list.Element),
		lru:      list.New(),
	}
}

// Get retrieves a value from the LRUCache based on the provided key.
//
// Parameters:
// - key: string representing the key of the value to retrieve.
// Return type:
// - interface{}: the value associated with the key, or nil if the key is not found.
func (c *LRUCache) Get(key string) interface{} {
	if elem, ok := c.data[key]; ok {
		entry := elem.Value.(*entry)

		if entry.expirationTime != nil {
			if entry.expirationTime.After(time.Now()) {
				c.lru.MoveToFront(elem)
				return entry.value
			} else {
				delete(c.data, key)
				c.lru.Remove(elem)
			}
		} else {
			c.lru.MoveToFront(elem)
			return entry.value
		}
	}

	return nil
}

// Set updates the LRUCache with the provided key, value, and expiration time.
//
// Parameters:
// - key: the key to be updated or added.
// - value: the value associated with the key.
// - expiration: the expiration time for the key-value pair.
// Set updates the LRUCache with the provided key, value, and expiration time.
//
// Parameters:
// - key: the key to be updated or added.
// - value: the value associated with the key.
// - expiration: the expiration time for the key-value pair.
func (c *LRUCache) Set(key string, value interface{}, expiration *time.Duration) {
	// Calculate the expiration time based on the input or set it to 1 hour if not provided
	var exp *time.Time

	if expiration != nil {
		expTime := time.Now().Add(*expiration)
		exp = &expTime
	} else {
		expTime := time.Now().Add(1 * time.Hour)
		exp = &expTime
	}

	// Update the value and expiration time if key already exists
	if elem, ok := c.data[key]; ok {
		c.lru.MoveToFront(elem)
		entry := elem.Value.(*entry)
		entry.value = value
		entry.expirationTime = exp
		return
	}

	// Remove the least recently used item if capacity is reached
	if c.capacity != nil {
		if c.lru.Len() >= int(*c.capacity) {
			elem := c.lru.Back()
			delete(c.data, elem.Value.(*entry).key)
			c.lru.Remove(elem)
		}
	}

	// Add the new key-value pair to the cache
	c.data[key] = c.lru.PushFront(&entry{
		key:            key,
		value:          value,
		expirationTime: exp,
	})
}
