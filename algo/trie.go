package algo

type Trie struct {
	// 支持多字符类型
	//node map[rune]*Trie
	// 只支持小写英文字符
	children [26]*Trie
	isTail   bool
	word     string
}

// 预处理共同前缀字符
func (t *Trie) InitPrefix(word string) {
	var temp *Trie = t
	for _, w := range []byte(word) {
		if temp.children[int(w-'a')] == nil {
			temp.children[int(w-'a')] = new(Trie)
		}
		temp = temp.children[int(w-'a')]
	}
	// 对每层节点设置标识
	temp.isTail = true
}

// 搜索共同前缀
func (t *Trie) SearchPrefix(word string) bool {
	node := t
	for _, ch := range []byte(word) {
		// 识别前缀中断
		if node.children[int(ch-'a')] == nil || !node.children[int(ch-'a')].isTail {
			return false
		}
		node = node.children[int(ch-'a')]
	}
	return true
}

func (t *Trie) InitDFA(word string) {
	var temp *Trie = t
	for _, w := range []byte(word) {
		if temp.children[int(w-'a')] == nil {
			temp.children[int(w-'a')] = new(Trie)
		}
		temp = temp.children[int(w-'a')]
	}
	// 对每层节点设置标识
	temp.word = word
}

func (t *Trie) Contain(word string) bool {
	node := t
	for _, ch := range []byte(word) {
		// 识别前缀中断
		if node.children[int(ch-'a')] == nil {
			return false
		}
		node = node.children[int(ch-'a')]
	}
	return node.word == word
}
