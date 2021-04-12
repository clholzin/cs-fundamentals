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

func (this *LRUCache) Update(key, value int) {
	if val, ok := this.Cache[key]; ok {
		val.Val = value
		valnext := val.Next
		valprev := val.Prev
		if valprev != nil {
			valprev.Next = valnext
			if valnext != nil {
				valnext.Prev = valprev
			}
			if valnext == nil {
				this.LastUsed.Tail = valprev
			}
		}

		head := this.LastUsed.Head
		if val != head {
			val.Next = head
			head.Prev = val
			val.Prev = nil
			this.LastUsed.Head = val
		}
	}
}
func (this *LRUCache) Add(key, value int) {
	head := this.LastUsed.Head
	val := &LinkNode{
		Key: key,
		Val: value,
	}
	val.Next = head
	if head != nil {
		head.Prev = val
	}
	this.LastUsed.Head = val
	if this.LastUsed.Tail == nil {
		this.LastUsed.Tail = head
	}
	this.Cache[key] = val
	this.Count++
}
func (this *LRUCache) Remove() {
	tail := this.LastUsed.Tail
	if tail == nil {
		tmp := this.LastUsed.Head
		this.LastUsed.Head = nil
		this.Count--
		delete(this.Cache, tmp.Key)
	} else {
		curr := tail
		prev := curr.Prev
		if prev != nil {
			this.LastUsed.Tail = prev
			this.LastUsed.Tail.Next = nil
		} else {
			this.LastUsed.Tail = nil
			this.LastUsed.Head = nil
		}
		this.Count--
		delete(this.Cache, curr.Key)
	}

}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.Cache[key]; ok {
		this.Update(key, val.Val)
		return val.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Cache[key]; !ok {
		if this.Count == this.Cap {
			this.Remove()
		}
		this.Add(key, value)
	} else {
		this.Update(key, value)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
