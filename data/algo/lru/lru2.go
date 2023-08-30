package lru

type LRUCache struct {
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func (this *LRUCache) AddToHead(node *DLinkedNode) {
	node.prev = nil
	node.next = this.head
	// 头节点非空
	if this.head != nil {
		// 更新原来头节点的指针
		this.head.prev = node
	}
	this.head = node
	// 第一次添加节点时，需要把队尾也指向自身
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}

func (this *LRUCache) Remove(node *DLinkedNode) {
	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == this.tail {
		this.tail = node.prev
		if node.prev != nil {
			node.prev.next = nil
		}
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	return
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, cache: make(map[int]*DLinkedNode)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.Remove(node)
		this.AddToHead(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.Remove(node)
		this.AddToHead(node)
		return
	}
	node := &DLinkedNode{key: key, value: value}
	this.AddToHead(node)
	this.cache[key] = node
	if len(this.cache) > this.capacity {
		delete(this.cache, this.tail.key)
		this.Remove(this.tail)
	}

}
