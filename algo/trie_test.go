package algo

import "testing"

// https://leetcode-cn.com/problems/longest-word-in-dictionary/
func TestTrie_SearchPrefix(t *testing.T) {
	trie := new(Trie)
	trie.InitPrefix("app")
	trie.InitPrefix("a")
	trie.InitPrefix("ap")
	trie.InitPrefix("appl")
	trie.InitPrefix("apply")
	t.Log(trie.SearchPrefix("apply"))
	t.Log(trie.SearchPrefix("app"))
}

// trie + dfa 敏感词过滤原理
func TestTrie_Contain(t *testing.T) {
	trie := new(Trie)
	trie.InitDFA("apple")
	trie.InitDFA("apply")
	t.Log(trie.Contain("apply"))
	t.Log(trie.Contain("app"))
}
