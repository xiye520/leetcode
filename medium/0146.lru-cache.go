package medium

import "container/list"

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}

type entry struct {
	key int
	Val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element),
		list:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.cache[key]; ok {
		this.list.MoveToFront(e)
		return e.Value.(entry).Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.cache[key]; ok {
		e.Value = entry{key, value}
		this.list.MoveToFront(e)
		return
	}

	listElement := this.list.PushFront(entry{key, value})
	this.cache[key] = listElement
	if len(this.cache) > this.cap {
		ele := this.list.Remove(this.list.Back())
		delete(this.cache, ele.(entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := NewLRUCache(cap);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
