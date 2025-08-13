package cache

import (
	"container/list"
	"orderService/internal/entity"
	"sync"
)

type LRUCache struct {
	sync.RWMutex
	items    map[string]*list.Element
	queue    *list.List
	Capacity int
}

type Item struct {
	Value entity.Order
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		items:    make(map[string]*list.Element),
		queue:    list.New(),
		Capacity: capacity,
	}
}

func (c *LRUCache) Set(key string, value entity.Order) {
	c.Lock()
	defer c.Unlock()

	if element, exists := c.items[key]; exists {
		element.Value.(*Item).Value = value
		c.queue.MoveToFront(element)
		return
	}

	if c.queue.Len() >= c.Capacity {
		c.removeOldest()
	}

	element := c.queue.PushFront(&Item{Value: value})
	c.items[key] = element
}

func (c *LRUCache) removeOldest() {
	if element := c.queue.Back(); element != nil {
		c.queue.Remove(element)
		delete(c.items, element.Value.(*Item).Value.OrderUID)
	}
}

func (c *LRUCache) Get(key string) (*entity.Order, bool) {
	c.RLock()
	defer c.RUnlock()

	element, exists := c.items[key]
	if !exists {
		return nil, false
	}
	c.queue.MoveToFront(element)
	return &element.Value.(*Item).Value, true
}
