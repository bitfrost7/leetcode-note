package design

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*LruNode
	head, tail *LruNode
}
type LruNode struct {
	key, value int
	prev, next *LruNode
}

func LRuCacheConstructor(capacity int) LRUCache {
	head, tail := &LruNode{}, &LruNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*LruNode),
		head:     head,
		tail:     tail,
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.cache[key]; ok {
		cache.moveToHead(node)
		return node.value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.cache[key]; ok {
		node.value = value
		cache.moveToHead(node)
	} else {
		node := &LruNode{
			key:   key,
			value: value,
		}
		cache.cache[key] = node
		cache.addNode(node)
		if cache.size > cache.capacity {
			node := cache.popTail()
			delete(cache.cache, node.key)
		}
	}
}

func (cache *LRUCache) addNode(node *LruNode) {
	node.prev = cache.head
	node.next = cache.head.next
	cache.head.next.prev = node
	cache.head.next = node
	cache.size++
}

func (cache *LRUCache) removeNode(node *LruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	cache.size--
}

func (cache *LRUCache) moveToHead(node *LruNode) {
	cache.removeNode(node)
	cache.addNode(node)
}

func (cache *LRUCache) popTail() *LruNode {
	node := cache.tail.prev
	cache.removeNode(node)
	return node
}
