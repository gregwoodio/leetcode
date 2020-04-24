package day24

type LRUCache struct {
	storage map[int]int
	usages  []int
	cap     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:     capacity,
		storage: make(map[int]int),
		usages:  make([]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.storage[key]; ok {
		this.moveKeyToStart(key)

		return val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.storage[key] = value
	this.moveKeyToStart(key)
}

func (this *LRUCache) moveKeyToStart(key int) {
	keyFound := false
	for i, usage := range this.usages {
		if usage == key {
			this.usages = append(append([]int{key}, this.usages[:i]...), this.usages[i+1:]...)
			keyFound = true
		}
	}

	if !keyFound {
		this.usages = append([]int{key}, this.usages...)
	}

	// remove the least used value from storage
	if len(this.usages) > this.cap {
		last := this.usages[this.cap]
		delete(this.storage, last)
		this.usages = this.usages[:this.cap]
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
