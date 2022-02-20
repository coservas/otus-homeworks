package hw04lrucache

import "sync"

type Key string

var m sync.Mutex

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	m.Lock()
	defer m.Unlock()

	item, ok := c.items[key]
	if ok {
		cItem := item.Value.(cacheItem)
		cItem.value = value
		item.Value = cItem
		c.queue.MoveToFront(item)

		return true
	}

	cItem := cacheItem{key, value}
	if c.capacity == c.queue.Len() {
		backItem := c.queue.Back()
		cItem := backItem.Value
		delete(c.items, cItem.(cacheItem).key)
		c.queue.Remove(backItem)
	}

	c.items[key] = c.queue.PushFront(cItem)

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(item)

	return item.Value.(cacheItem).value, true
}

func (c *lruCache) Clear() {
	m.Lock()
	defer m.Unlock()

	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
