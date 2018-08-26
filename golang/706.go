type MyHashMap struct {
    Key   []int
    Value []int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    ret := MyHashMap{}
    ret.Key = make([]int, 0)
    ret.Value = make([]int, 0)
    return ret
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    for i, k := range this.Key {
        if k == key {
            this.Value[i] = value
            return
        }
    }
    this.Key = append(this.Key, key)
    this.Value = append(this.Value, value)
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    for i, k := range this.Key {
        if k == key {
            return this.Value[i] 
        }
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    for i, k := range this.Key {
        if k == key {
            this.Key = append(this.Key[:i], this.Key[i+1:]...)
            this.Value = append(this.Value[:i], this.Value[i+1:]...)
            return
        }
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
 