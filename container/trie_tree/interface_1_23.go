//go:build go1.23
// +build go1.23

// Package trie_tree provides a Trie (prefix tree) data structure implementation.
// A Trie is a tree-like data structure that stores strings efficiently and
// supports fast prefix-based operations.

package trie_tree

import "iter"

// Interface defines the operations for a Trie data structure.
// A Trie is optimal for storing and searching strings with common prefixes.
type Interface interface {
	// Insert adds a word to the trie.
	Insert(word string)

	// Search returns true if the word exists in the trie.
	Search(word string) bool

	// StartsWith returns true if there are any words in the trie that start with the given prefix.
	StartsWith(prefix string) bool

	// Delete removes a word from the trie and returns true if the word was found and removed.
	Delete(word string) bool

	// Len returns the number of words stored in the trie.
	Len() int

	// Clear removes all words from the trie.
	Clear()

	// GetAllWords returns a slice of all words stored in the trie in lexicographical order.
	GetAllWords() []string

	// GetWordsWithPrefix returns a slice of all words that start with the given prefix
	// in lexicographical order.
	GetWordsWithPrefix(prefix string) []string

	// WordSeq returns an iterator over all words in the trie in lexicographical order.
	WordSeq() iter.Seq[string]

	// PrefixSeq returns an iterator over all words that start with the given prefix
	// in lexicographical order.
	PrefixSeq(prefix string) iter.Seq[string]
}
