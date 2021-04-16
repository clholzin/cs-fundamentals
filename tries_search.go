package fundamentals

// https://leetcode.com/problems/implement-trie-prefix-tree/
/*
Runtime: 80 ms, faster than 29.38% of Go online submissions for Implement Trie (Prefix Tree).
Memory Usage: 22.8 MB, less than 5.15% of Go online submissions for Implement Trie (Prefix Tree).
*/
type Trie struct {
	Words map[string]interface{}
}

/** Initialize your data structure here. */
func ConstructorLRU() Trie {
	words := make(map[string]interface{})
	return Trie{Words: words}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.recurAdd(this.Words, word, 0)
}

func (this *Trie) recurAdd(words map[string]interface{}, word string, index int) {
	if index == len(word) {
		if _, ok := words["-"+word]; !ok {
			words["-"+word] = make(map[string]interface{})
		}
		return
	}
	if _, ok := words[string(word[index])]; !ok {
		words[string(word[index])] = make(map[string]interface{})
	}
	this.recurAdd(words[string(word[index])].(map[string]interface{}), word, index+1)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	ok := this.recurSearch(this.Words, word, 0)
	return ok
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	ok := this.recurPrefix(this.Words, prefix, 0)
	return ok
}

func (this *Trie) recurPrefix(words interface{}, word string, index int) bool {
	if index == len(word) {
		return true
	}
	w := words.(map[string]interface{})
	if _, ok := w[string(word[index])]; ok {
		return this.recurPrefix(w[string(word[index])], word, index+1)
	}
	return false
}

func (this *Trie) recurSearch(words interface{}, word string, index int) bool {
	if index == len(word) {
		w := words.(map[string]interface{})
		_, ok := w["-"+word]
		return ok
	}
	w := words.(map[string]interface{})
	if _, ok := w[string(word[index])]; ok {
		return this.recurSearch(w[string(word[index])], word, index+1)
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
