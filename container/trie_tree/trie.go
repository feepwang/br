// Package trie_tree provides a Trie (prefix tree) data structure implementation.
// This file implements the Interface using a standard Trie algorithm.

package trie_tree

import (
	"sort"
)

// trieNode represents a node in the Trie tree.
type trieNode struct {
	children map[rune]*trieNode // children nodes mapped by character
	isEnd    bool               // true if this node represents the end of a word
}

// newTrieNode creates a new trie node.
func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
		isEnd:    false,
	}
}

// Trie implements the Interface using a standard Trie data structure.
// It uses a tree of nodes where each edge represents a character.
type Trie struct {
	root *trieNode
	size int // number of words stored
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
		size: 0,
	}
}

// Insert adds a word to the trie.
func (t *Trie) Insert(word string) {
	if word == "" {
		return
	}

	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = newTrieNode()
		}
		node = node.children[char]
	}

	// Mark the end of the word
	if !node.isEnd {
		node.isEnd = true
		t.size++
	}
}

// Search returns true if the word exists in the trie.
func (t *Trie) Search(word string) bool {
	if word == "" {
		return false
	}

	node := t.findNode(word)
	return node != nil && node.isEnd
}

// StartsWith returns true if there are any words in the trie that start with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return t.size > 0
	}

	return t.findNode(prefix) != nil
}

// Delete removes a word from the trie and returns true if the word was found and removed.
func (t *Trie) Delete(word string) bool {
	if word == "" {
		return false
	}

	// First check if the word exists
	node := t.findNode(word)
	if node == nil || !node.isEnd {
		return false
	}

	// Word exists, so remove it
	t.deleteHelper(t.root, word, 0)
	return true
}

// deleteHelper is a recursive helper function for deletion.
func (t *Trie) deleteHelper(node *trieNode, word string, index int) bool {
	chars := []rune(word)
	
	if index == len(chars) {
		// We've reached the end of the word
		if !node.isEnd {
			return false // Word doesn't exist
		}
		node.isEnd = false
		t.size--
		// Return true if current node has no children (can be deleted)
		return len(node.children) == 0
	}

	char := chars[index]
	childNode, exists := node.children[char]
	if !exists {
		return false // Word doesn't exist
	}

	shouldDeleteChild := t.deleteHelper(childNode, word, index+1)

	if shouldDeleteChild {
		delete(node.children, char)
		// Return true if current node is not end of another word and has no children
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

// Len returns the number of words stored in the trie.
func (t *Trie) Len() int {
	return t.size
}

// Clear removes all words from the trie.
func (t *Trie) Clear() {
	t.root = newTrieNode()
	t.size = 0
}

// GetAllWords returns a slice of all words stored in the trie in lexicographical order.
func (t *Trie) GetAllWords() []string {
	var words []string
	t.collectWords(t.root, "", &words)
	return words
}

// GetWordsWithPrefix returns a slice of all words that start with the given prefix
// in lexicographical order.
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	var words []string
	
	if prefix == "" {
		return t.GetAllWords()
	}

	prefixNode := t.findNode(prefix)
	if prefixNode == nil {
		return words // Return empty slice if prefix doesn't exist
	}

	// Collect all words that start with the prefix
	t.collectWords(prefixNode, prefix, &words)
	
	return words
}

// findNode traverses the trie to find the node representing the given string.
// Returns nil if the string is not found.
func (t *Trie) findNode(str string) *trieNode {
	node := t.root
	for _, char := range str {
		if child, exists := node.children[char]; exists {
			node = child
		} else {
			return nil
		}
	}
	return node
}

// collectWords performs a depth-first search to collect all words from a given node.
func (t *Trie) collectWords(node *trieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}

	// Get all children characters and sort them for consistent ordering
	var chars []rune
	for char := range node.children {
		chars = append(chars, char)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	// Recursively collect words from children
	for _, char := range chars {
		child := node.children[char]
		t.collectWords(child, prefix+string(char), words)
	}
}