package main

type TrieNode struct {
	n [26]*TrieNode
}
type Trie struct {
	root *TrieNode	
}

func NewTrie() *Trie{
	return &Trie {
		root: &TrieNode{
		},
	}
}

func (t *Trie) Insert(word string){
	node := t.root
	for _, ch := range word {
		index := ch - 'a'
		if node.n[index] == nil {
			node.n[index] = &TrieNode{}
		}
		node = node.n[index]
	}
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		index := ch - 'a'
		if node.n[index] == nil {
			return false
		}
		node = node.n[index]
	}
	return true
}