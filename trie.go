// Divyansh Khanna 2015

// trie data structure in GoLang

package main

import "fmt"

// global variable
const ALPHABETSIZE int = 26

// trie node
type trie_node struct {
	value    int
	children [ALPHABETSIZE]*trie_node
}

// trie ADT
type trie struct {
	count int
	root  *trie_node
}

// returns new trie node
func getNode() *trie_node {
	node := new(trie_node)
	node.value = 0
	for i := 0; i < ALPHABETSIZE; i++ {
		node.children[i] = nil
	}
	return node
}

// Initializes trie (root is dummy node)
func initial(ptrie *trie) {
	ptrie.count = 0
	ptrie.root = getNode()
}

// If not present, inserts key into trie
// If the key is prefix of trie node, just marks leaf node
func insert(ptrie *trie, key string) {
	length := len(key)
	pcrawl := new(trie_node)
	pcrawl = ptrie.root
	for level := 0; level < length; level++ {
		index := key[level] - 97
		if pcrawl.children[index] == nil {
			pcrawl.children[index] = getNode()
		}
		pcrawl = pcrawl.children[index]
	}
	// mark last node as leaf
	pcrawl.value = 1
}

// Bool to Int
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Returns non zero if key present in trie
func search(ptrie *trie, key string) int {
	length := len(key)
	pcrawl := new(trie_node)
	pcrawl = ptrie.root
	for level := 0; level < length; level++ {
		index := key[level] - 97
		if pcrawl.children[index] == nil {
			return 0
		}
		pcrawl = pcrawl.children[index]
	}
	return Btoi(pcrawl != nil && pcrawl.value != 0)
}

// Driver

func main() {
	keys := []string{"the", "a", "there", "answer", "any", "by", "bye", "their"}
	output := []string{"not present in trie", "present in trie"}

	ptrie := new(trie) // new returns a pointer itself!! good boy!!
	initial(ptrie)

	for i := 0; i < len(keys); i++ {
		insert(ptrie, keys[i])
	}

	// Search for different keys
	fmt.Println("the", " --- ", output[search(ptrie, "the")])
	fmt.Println("these", " --- ", output[search(ptrie, "these")])
	fmt.Println("their", " --- ", output[search(ptrie, "their")])
	fmt.Println("thaw", " --- ", output[search(ptrie, "thaw")])
}
