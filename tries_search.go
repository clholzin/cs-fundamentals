package fundamentals

// https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
	Words map[string]interface{}
	Size  int
}

/** Initialize your data structure here. */
func Constructor() Trie {
	words := make(map[string]interface{})
	return Trie{Words: words, Size: 0}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.Size += len(word)
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
	//fmt.Println(words)
	this.recurAdd(words[string(word[index])].(map[string]interface{}), word, index+1)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	ok := this.recurSearch(this.Words, word, 0)
	//fmt.Println("search",this.Words,word,ok)
	return ok
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	ok := this.recurPrefix(this.Words, prefix, 0)
	//fmt.Println("prefix",this.Words,prefix,ok)
	return ok
}

func (this *Trie) recurPrefix(words interface{}, word string, index int) bool {
	if index == len(word) {
		//fmt.Println("index == len(word)",word,index)
		return true
	}
	var w map[string]interface{} = words.(map[string]interface{})
	if _, ok := w[string(word[index])]; ok {
		return this.recurPrefix(w[string(word[index])], word, index+1)
	}
	return false
}

func (this *Trie) recurSearch(words interface{}, word string, index int) bool {
	if index == len(word) {
		//fmt.Println("index == len(word)",word,index)
		w := words.(map[string]interface{})
		if _, ok := w["-"+word]; ok {
			return true
		}
		return false
	}
	var w map[string]interface{} = words.(map[string]interface{})
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
