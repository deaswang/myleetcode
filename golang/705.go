type MyHashSet struct {
    Values map[int]bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    ret := MyHashSet{}
    ret.Values = make(map[int]bool)
    return ret
}


func (this *MyHashSet) Add(key int) {
    if _, ok := this.Values[key]; !ok {
        this.Values[key] = true
    }
}


func (this *MyHashSet) Remove(key int)  {
    if _, ok := this.Values[key]; ok {
        delete(this.Values, key)
    }
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    _, ok := this.Values[key]
    return ok
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 