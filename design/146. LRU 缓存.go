package design

// https://leetcode.cn/problems/lru-cache/description/
// 使用双向链表维护缓存的上一次使用时间：
//
// 约定：链表正方向（从头部到尾部）节点按照使用时间排序——越早使用（即久未使用）的节点，越靠近链表尾部
// 维护：每使用一次缓存，就将该缓存对应的链表节点移动到链表头部；缓存淘汰时，只需要删除尾部节点即可
// 增加一个map，记录key到链表节点的映射关系; 解决如果只使用双向链表，每次判断key是否存在时，都要遍历链表

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
