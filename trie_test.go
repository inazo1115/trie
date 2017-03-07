package trie

import "testing"

func TestTrie(t *testing.T) {

	trie := NewTrie()

	trie.Insert([]rune("aaa"))
	trie.Insert([]rune("abc"))
	trie.Insert([]rune("abd"))
	trie.Insert([]rune("あいう"))
	trie.Insert([]rune("あいえ"))

	trie.Dump(0)
	// # Node
	// - a
	//     # Node
	//     - b
	//         # Node
	//         - c
	//         - d
	//     - a
	//         # Node
	//         - a
	// - あ
	//     # Node
	//     - い
	//         # Node
	//         - う
	//         - え

	if _, ok := trie.Find([]rune("abc")); !ok {
		t.Errorf("error Find")
	}
	if _, ok := trie.Find([]rune("あいう")); !ok {
		t.Errorf("error Find")
	}
	node, ok := trie.Find([]rune("ab"))
	if !ok {
		t.Errorf("error Find")
	}
	if _, ok := node.Find([]rune("d")); !ok {
		t.Errorf("error Find")
	}
}
