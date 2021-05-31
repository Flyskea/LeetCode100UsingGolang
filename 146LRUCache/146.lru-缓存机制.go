/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */
package main

// @lc code=start
type doubleList struct {
	prev, next *doubleList
	key, val   int
}
type LRUCache struct {
	head, tail *doubleList
	keys       map[int]*doubleList
	capacity   int
}

func (lru *LRUCache) add(node *doubleList) {
	node.prev = nil
	node.next = lru.head
	if lru.head != nil {
		lru.head.prev = node
	}
	lru.head = node
	if lru.tail == nil {
		lru.tail = node
		lru.tail.next = nil
	}
}

func (lru *LRUCache) remove(node *doubleList) {
	if node == lru.head {
		lru.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == lru.tail {
		lru.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func Constructor(capacity int) LRUCache {
	return LRUCache{keys: make(map[int]*doubleList), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.keys[key]; ok {
		this.remove(node)
		this.add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.keys[key]; ok {
		node.val = value
		this.remove(node)
		this.add(node)
		return
	} else {
		node = &doubleList{key: key, val: value}
		this.keys[key] = node
		this.add(node)
	}
	if len(this.keys) > this.capacity {
		delete(this.keys, this.tail.key)
		this.remove(this.tail)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
