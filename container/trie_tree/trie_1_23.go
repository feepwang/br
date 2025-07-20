//go:build go1.23
// +build go1.23

// Package trie_tree provides go1.23-specific methods for Trie.
// This file adds iter.Seq related methods for Interface.

package trie_tree

import (
	"iter"
	"sort"
)

// WordSeq returns an iterator for all words in the trie in lexicographical order (go1.23).
// Uses efficient depth-first traversal without pre-allocating all words.
func (t *Trie) WordSeq() iter.Seq[string] {
	return func(yield func(string) bool) {
		collectWordsIterative(t.root, "", yield)
	}
}

// PrefixSeq returns an iterator for all words that start with the given prefix
// in lexicographical order (go1.23).
// Uses efficient depth-first traversal without pre-allocating all words.
func (t *Trie) PrefixSeq(prefix string) iter.Seq[string] {
	return func(yield func(string) bool) {
		node := t.findNode(prefix)
		if node != nil {
			collectWordsIterative(node, prefix, yield)
		}
	}
}

// collectWordsIterative performs depth-first search to iterate over all words from a given node.
// It yields words in lexicographical order and stops early if yield returns false.
// Returns false if iteration should stop (early termination requested).
func collectWordsIterative(node *trieNode, prefix string, yield func(string) bool) bool {
	if node == nil {
		return true // Continue iteration
	}

	// If this node represents the end of a word, yield it
	if node.isEnd {
		if !yield(prefix) {
			return false // Stop iteration if yield returns false
		}
	}

	// Get all children characters and sort them for consistent lexicographical ordering
	var chars []rune
	for char := range node.children {
		chars = append(chars, char)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	// Recursively iterate through children in sorted order
	for _, char := range chars {
		child := node.children[char]
		if !collectWordsIterative(child, prefix+string(char), yield) {
			return false // Propagate early termination
		}
	}

	return true // Continue iteration
}
