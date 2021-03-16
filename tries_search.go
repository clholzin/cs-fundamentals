package fundamentals

// https://leetcode.com/problems/implement-trie-prefix-tree/
type Trie struct {
	Words map[byte]interface{}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Words: make(map[byte]interface{})}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.recurAdd(this.Words, word, 0)
}

func (this *Trie) recurAdd(words interface{}, word string, index int) {
	if index == len(word) {
		return
	}
	var w map[byte]interface{} = words.(map[byte]interface{})
	if _, ok := w[word[index]]; !ok {
		w = make(map[byte]interface{})
	}
	this.recurAdd(w[word[index]], word, index+1)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.recurSearch(this.Words, word, 0)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.recurSearch(this.Words, prefix, 0)
}

func (this *Trie) recurSearch(words interface{}, word string, index int) bool {
	if index == len(word) {
		return true
	}
	var w map[byte]interface{} = words.(map[byte]interface{})
	if _, ok := w[word[index]]; ok {
		return this.recurSearch(w[word[index]], word, index+1)
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
