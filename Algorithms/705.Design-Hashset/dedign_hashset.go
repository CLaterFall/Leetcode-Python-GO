type MyHashSet struct {
	table []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{table: make([]bool, 1000001)}
}

func (this *MyHashSet) Add(key int) {
	this.table[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.table[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.table[key]
}
