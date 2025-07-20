//go:build go1.23
// +build go1.23

package trie_tree

import (
	"slices"
	"testing"
)

func TestTrieWordSeq(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "banana", "band"}

	for _, word := range words {
		trie.Insert(word)
	}

	// Collect words using iterator
	var collected []string
	for word := range trie.WordSeq() {
		collected = append(collected, word)
	}

	// Compare with GetAllWords (should be identical)
	expected := trie.GetAllWords()
	if !slices.Equal(collected, expected) {
		t.Errorf("WordSeq() = %v, want %v", collected, expected)
	}

	// Verify lexicographical order
	if !slices.IsSorted(collected) {
		t.Errorf("WordSeq() result is not sorted: %v", collected)
	}
}

func TestTriePrefixSeq(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "banana", "band"}

	for _, word := range words {
		trie.Insert(word)
	}

	// Test prefix "app"
	var collected []string
	for word := range trie.PrefixSeq("app") {
		collected = append(collected, word)
	}

	// Compare with GetWordsWithPrefix (should be identical)
	expected := trie.GetWordsWithPrefix("app")
	if !slices.Equal(collected, expected) {
		t.Errorf("PrefixSeq(\"app\") = %v, want %v", collected, expected)
	}

	// Verify lexicographical order
	if !slices.IsSorted(collected) {
		t.Errorf("PrefixSeq(\"app\") result is not sorted: %v", collected)
	}
}

func TestTriePrefixSeqNonExistent(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application"}

	for _, word := range words {
		trie.Insert(word)
	}

	// Test non-existent prefix
	var collected []string
	for word := range trie.PrefixSeq("xyz") {
		collected = append(collected, word)
	}

	if len(collected) != 0 {
		t.Errorf("PrefixSeq(\"xyz\") = %v, want empty slice", collected)
	}
}

func TestTrieWordSeqEmpty(t *testing.T) {
	trie := NewTrie()

	var collected []string
	for word := range trie.WordSeq() {
		collected = append(collected, word)
	}

	if len(collected) != 0 {
		t.Errorf("WordSeq() on empty trie = %v, want empty slice", collected)
	}
}

func TestTrieIteratorEarlyStop(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "banana", "band"}

	for _, word := range words {
		trie.Insert(word)
	}

	// Stop after collecting 3 words
	var collected []string
	for word := range trie.WordSeq() {
		collected = append(collected, word)
		if len(collected) >= 3 {
			break
		}
	}

	if len(collected) != 3 {
		t.Errorf("Early stop failed: got %d words, want 3", len(collected))
	}

	// Verify they are still in lexicographical order
	if !slices.IsSorted(collected) {
		t.Errorf("Early stopped WordSeq() result is not sorted: %v", collected)
	}
}

func TestTrieIteratorUnicodeSupport(t *testing.T) {
	trie := NewTrie()
	words := []string{"你好", "你", "世界", "测试"}

	for _, word := range words {
		trie.Insert(word)
	}

	// Test WordSeq with Unicode
	var collected []string
	for word := range trie.WordSeq() {
		collected = append(collected, word)
	}

	expected := trie.GetAllWords()
	if !slices.Equal(collected, expected) {
		t.Errorf("WordSeq() with Unicode = %v, want %v", collected, expected)
	}

	// Test PrefixSeq with Unicode
	var prefixCollected []string
	for word := range trie.PrefixSeq("你") {
		prefixCollected = append(prefixCollected, word)
	}

	expectedPrefix := trie.GetWordsWithPrefix("你")
	if !slices.Equal(prefixCollected, expectedPrefix) {
		t.Errorf("PrefixSeq(\"你\") = %v, want %v", prefixCollected, expectedPrefix)
	}
}
