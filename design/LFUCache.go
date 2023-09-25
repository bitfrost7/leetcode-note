package design

// LFUCache : least frequently used cache.
// the size is the current size of cache, capacity is the max size of cache;
// minFreq is the min freq of cache.
// freqMap is a map of freq and LfuList.
// keyMap is a map of key and LfuNode.
type LFUCache struct {
	size     int
	capacity int
	minFreq  int
	freqMap  map[int]*LfuList
	keyMap   map[int]*LfuNode
}

// LfuList : a Two-way queues of LfuNode.
// The head node is a sentinel node that does not store data and is convenient for delete or append elements.
type LfuList struct {
	head, tail *LfuNode
}

// LfuNode : a node of LfuList. The freq field stores the frequency of element access
type LfuNode struct {
	key, value, freq int
	prev, next       *LfuNode
}

// LfuCacheConstructor : initialize a LFUCache.
func LfuCacheConstructor(capacity int) LFUCache {
	return LFUCache{
		size:     0,
		capacity: capacity,
		minFreq:  0,
		freqMap:  map[int]*LfuList{},
		keyMap:   map[int]*LfuNode{},
	}
}

// Get : get the value of key in cache. increase the frequency of the key. If the key does not exist, return -1.
func (cache *LFUCache) Get(key int) int {
	if node, ok := cache.keyMap[key]; ok {
		node = cache.IncrFreq(node)
		return node.value
	}
	return -1
}

// Put : put the key and value in cache. If the key already exists, update the value.
// If the key does not exist, create a new node and add it to the cache. If the cache is full, delete the node with the lowest frequency.
func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}
	if node, ok := cache.keyMap[key]; ok {
		node.value = value
		node = cache.IncrFreq(node)
	} else {
		node := &LfuNode{
			key:   key,
			value: value,
			freq:  1,
		}
		cache.keyMap[key] = node
		cache.AddNode(node)
		cache.size++
		if cache.size > cache.capacity {
			cache.RemoveMinFreqNode()
			cache.size--
		}
		cache.minFreq = 1
	}
}

// AddNode : add a node to the cache. If the frequency of the node is not in the cache, create a new LfuList and add it to the cache.
func (cache *LFUCache) AddNode(node *LfuNode) {
	if list, ok := cache.freqMap[node.freq]; ok {
		list.addNode(node)
	} else {
		list := &LfuList{
			head: &LfuNode{},
			tail: &LfuNode{},
		}
		list.head.next = list.tail
		list.tail.prev = list.head
		list.addNode(node)
		cache.freqMap[node.freq] = list
	}
}

// RemoveMinFreqNode : remove the node with the lowest frequency from the cache.
func (cache *LFUCache) RemoveMinFreqNode() {
	if list, ok := cache.freqMap[cache.minFreq]; ok {
		node := list.tail.prev
		list.removeNode(node)
		if list.head.next == list.tail {
			delete(cache.freqMap, node.freq)
		}
		delete(cache.keyMap, node.key)
	}
}

// IncrFreq : increase the frequency of the node. If the frequency of the node is the lowest frequency, update the minFreq.
func (cache *LFUCache) IncrFreq(node *LfuNode) *LfuNode {
	if list, ok := cache.freqMap[node.freq]; ok {
		list.removeNode(node)
		if list.head.next == list.tail {
			delete(cache.freqMap, node.freq)
			if node.freq == cache.minFreq {
				cache.minFreq++
			}
		}
		node.freq++
		cache.AddNode(node)
		return node
	}
	return nil
}

// addNode : add a node to the LfuList.
func (l *LfuList) addNode(node *LfuNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

// removeNode : remove a node from the LfuList.
func (l *LfuList) removeNode(node *LfuNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
