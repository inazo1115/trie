package trie

import (
	"fmt"
	"strings"
)

type Node struct {
	children map[rune]*Node
}

func NewTrie() *Node {
	return &Node{}
}

func (n *Node) Insert(key []rune) {
	if len(key) == 0 {
		return
	}
	if n.children == nil {
		n.children = make(map[rune]*Node)
	}
	if _, ok := n.children[key[0]]; !ok {
		n.children[key[0]] = &Node{}
	}
	n.children[key[0]].Insert(key[1:])
}

func (n *Node) Find(key []rune) (bool, *Node) {
	if len(key) == 0 {
		return true, n
	}
	if n.children == nil || len(n.children) == 0 {
		return false, nil
	}
	child, ok := n.children[key[0]]
	if !ok {
		return false, nil
	}
	return child.Find(key[1:])
}

func (n *Node) Dump(level int) {
	if n.children == nil || len(n.children) == 0 {
		return
	}
	indent := strings.Repeat("    ", level)
	fmt.Printf("%s# Node\n", indent)
	for k, v := range n.children {
		fmt.Printf("%s- %s\n", indent, string(k))
		v.Dump(level + 1)
	}
}
