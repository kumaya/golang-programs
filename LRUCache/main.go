package LRUCache

import (
    "container/list"
)

type LRUCache struct {
	// maximum capacity of the cache
	capacity int
	list     *list.List
	items    map[interface{}]*list.Element
}

// a key can be any value
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

// New creates a new LRUCache
func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		items:    make(map[interface{}]*list.Element),
	}
}

// Set creates a value in LRUCache
func (c *LRUCache) Set(key Key, value interface{}) bool {
	item, exists := c.items[key]
    if exists {
        c.list.MoveToFront(item)
        item.Value.(*entry).value = value
        return true
    }
    if c.capacity <= c.list.Len() {
        c.prune()
    }
    newItem := c.list.PushFront(&entry{key, value})
	c.items[key] = newItem
	return true
}

func (c *LRUCache) prune() {
	tail := c.list.Back()
	if tail == nil {
		return
	}
	c.list.Remove(tail)
	kvPair := tail.Value.(*entry)
	delete(c.items, kvPair.key)
}

// Get look up a keys value from LRUCache
func (c *LRUCache) Get(key Key) (interface{}, bool) {
	item, exists := c.items[key]
	if !exists {
		return nil, false
	}
	c.list.MoveToFront(item)
	return item.Value.(*entry).value, true
}
