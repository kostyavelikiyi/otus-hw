package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) purge(key Key) {
	element := lc.queue.Back()
	if element != nil {
		lc.queue.Remove(element)
		delete(lc.items, element.Value.(cacheItem).key)
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	element, exists := lc.items[key]
	ci := cacheItem{
		key:   key,
		value: value,
	}
	if exists {
		lc.queue.MoveToFront(element)
		(*element).Value = ci
		return true
	}

	if lc.queue.Len() == lc.capacity {
		lc.purge(key)
	}

	item := lc.queue.PushFront(ci)
	lc.items[key] = item

	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	element, exists := lc.items[key]
	if !exists {
		return nil, false
	}
	lc.queue.MoveToFront(element)
	return element.Value.(cacheItem).value, true
}

func (lc *lruCache) Clear() {
	return
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
