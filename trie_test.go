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

	if found, _ := trie.Find([]rune("abc")); !found {
		t.Errorf("error Find")
	}
	if found, _ := trie.Find([]rune("あいう")); !found {
		t.Errorf("error Find")
	}
	found, node := trie.Find([]rune("ab"))
	if !found {
		t.Errorf("error Find")
	}
	if found, _ := node.Find([]rune("d")); !found {
		t.Errorf("error Find")
	}
}
