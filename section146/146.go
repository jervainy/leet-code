package main

type Node struct {
	value int
	age int
}

type LRUCache struct {
	data map[int]Node
	length int
	capacity int
	count int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		data: make(map[int]Node),
		length: 0,
		capacity: capacity,
		count: 0,
	}
}


func (this *LRUCache) Get(key int) int {
	count := this.count
	this.count += 1
	if obj, ok := this.data[key]; ok {
		obj.age = count
		this.data[key] = obj
		return obj.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	obj := Node{ value: value, age: this.count }
	if this.length >= this.capacity { // 满了
		if _, ok := this.data[key]; !ok {
			minAge, minKey := -1, 0
			for i, node := range this.data {
				if minAge == -1 || node.age < minAge {
					minAge = node.age
					minKey = i
				}
			}
			delete(this.data, minKey)
		}
	} else {
		this.length += 1
	}
	this.count += 1
	this.data[key] = obj
}

func main() {
	// TODO 待debug
	obj := Constructor(2)
	obj.Get(2)
	obj.Put(2, 6)
	obj.Get(1)
	obj.Put(1, 5)
	obj.Put(1, 2)
	obj.Get(1)
	obj.Get(2)
	println(1)
}
