package fundamentals

type Link struct {
	Head *LinkNode
	Tail *LinkNode
}

type LinkNode struct {
	Key  int
	Val  int
	Next *LinkNode
	Prev *LinkNode
}

type LRUCache struct {
	Cache    map[int]*LinkNode
	LastUsed *Link
	Count    int
	Cap      int
}

func NewLink() *Link {
	return &Link{}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache:    make(map[int]*LinkNode),
		LastUsed: NewLink(),
		Cap:      capacity,
	}
}

func (lru *LRUCache) Update(key, value int) {
	if val, ok := lru.Cache[key]; ok {
		val.Val = value
		valnext := val.Next
		valprev := val.Prev
		if valprev != nil {
			valprev.Next = valnext
			if valnext != nil {
				valnext.Prev = valprev
			}
			if valnext == nil {
				lru.LastUsed.Tail = valprev
			}
		}

		head := lru.LastUsed.Head
		if val != head {
			val.Next = head
			head.Prev = val
			val.Prev = nil
			lru.LastUsed.Head = val
		}
	}
}
func (lru *LRUCache) Add(key, value int) {
	head := lru.LastUsed.Head
	val := &LinkNode{
		Key: key,
		Val: value,
	}
	val.Next = head
	if head != nil {
		head.Prev = val
	}
	lru.LastUsed.Head = val
	if lru.LastUsed.Tail == nil {
		lru.LastUsed.Tail = head
	}
	lru.Cache[key] = val
	lru.Count++
}
func (lru *LRUCache) Remove() {
	tail := lru.LastUsed.Tail
	if tail == nil {
		tmp := lru.LastUsed.Head
		lru.LastUsed.Head = nil
		lru.Count--
		delete(lru.Cache, tmp.Key)
	} else {
		curr := tail
		prev := curr.Prev
		if prev != nil {
			lru.LastUsed.Tail = prev
			lru.LastUsed.Tail.Next = nil
		} else {
			lru.LastUsed.Tail = nil
			lru.LastUsed.Head = nil
		}
		lru.Count--
		delete(lru.Cache, curr.Key)
	}

}

func (lru *LRUCache) Get(key int) int {
	if val, ok := lru.Cache[key]; ok {
		lru.Update(key, val.Val)
		return val.Val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.Cache[key]; !ok {
		if lru.Count == this.Cap {
			lru.Remove()
		}
		lru.Add(key, value)
	} else {
		lru.Update(key, value)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/* example set
["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
*/
