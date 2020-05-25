/*
 * @Time     : 2020/5/25 14:14
 * @Author   : cancan
 * @File     : 146.go
 * @Function : LRU缓存机制
 */

/*
 * Question:
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 *
 * Follow up:
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
 * Example:
 * LRUCache cache = new LRUCache( 2 缓存容量 );
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 */

package QuestionBank

type LRUCache struct {
	capacity int
	order    []int
	data     map[int]int
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		order:    make([]int, 0, capacity),
		data:     make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	if ret, ok := this.data[key]; ok {
		this.delOrderItem(key)
		this.order = append(this.order, key)
		return ret
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		this.delOrderItem(key)
	}
	this.data[key] = value
	this.order = append(this.order, key)

	l := len(this.order)
	if l > this.capacity {
		d := l - this.capacity
		for _, k := range this.order[:d] {
			delete(this.data, k)
		}
		this.order = this.order[d:]
	}
}

func (this *LRUCache) delOrderItem(key int) {
	for idx, k := range this.order {
		if k == key {
			this.order = append(this.order[:idx], this.order[idx+1:]...)
			break
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
