package trie_tree

import (
	"reflect"
	"testing"
)

func TestTrieBasic(t *testing.T) {
	trie := NewTrie()

	// Test empty trie
	if trie.Len() != 0 {
		t.Errorf("Expected length 0, got %d", trie.Len())
	}

	// Test Search on empty trie
	if trie.Search("hello") {
		t.Error("Expected false when searching in empty trie")
	}

	// Test StartsWith on empty trie
	if trie.StartsWith("he") {
		t.Error("Expected false when checking prefix in empty trie")
	}

	// Test GetAllWords on empty trie
	words := trie.GetAllWords()
	if len(words) != 0 {
		t.Errorf("Expected empty slice, got %v", words)
	}
}

func TestTrieInsertAndSearch(t *testing.T) {
	trie := NewTrie()

	// Insert words
	words := []string{"hello", "world", "help", "he", "her", "hero"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test length
	if trie.Len() != len(words) {
		t.Errorf("Expected length %d, got %d", len(words), trie.Len())
	}

	// Test searching for existing words
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find word '%s'", word)
		}
	}

	// Test searching for non-existing words
	nonExistingWords := []string{"hel", "helping", "wor", "heroes"}
	for _, word := range nonExistingWords {
		if trie.Search(word) {
			t.Errorf("Expected not to find word '%s'", word)
		}
	}
}

func TestTrieStartsWith(t *testing.T) {
	trie := NewTrie()

	// Insert words
	words := []string{"hello", "help", "hero", "world"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test existing prefixes
	prefixes := []string{"he", "hel", "help", "hero", "w", "wo", "wor", "world"}
	for _, prefix := range prefixes {
		if !trie.StartsWith(prefix) {
			t.Errorf("Expected to find prefix '%s'", prefix)
		}
	}

	// Test non-existing prefixes
	nonExistingPrefixes := []string{"hi", "hal", "word", "hello!"}
	for _, prefix := range nonExistingPrefixes {
		if trie.StartsWith(prefix) {
			t.Errorf("Expected not to find prefix '%s'", prefix)
		}
	}
}

func TestTrieDelete(t *testing.T) {
	trie := NewTrie()

	// Insert words
	words := []string{"hello", "help", "hero", "her", "he"}
	for _, word := range words {
		trie.Insert(word)
	}

	initialLen := trie.Len()

	// Delete existing word
	if !trie.Delete("hello") {
		t.Error("Expected to successfully delete 'hello'")
	}

	if trie.Len() != initialLen-1 {
		t.Errorf("Expected length %d after deletion, got %d", initialLen-1, trie.Len())
	}

	if trie.Search("hello") {
		t.Error("Expected 'hello' to be deleted")
	}

	// Ensure other words still exist
	remainingWords := []string{"help", "hero", "her", "he"}
	for _, word := range remainingWords {
		if !trie.Search(word) {
			t.Errorf("Expected word '%s' to still exist after deleting 'hello'", word)
		}
	}

	// Delete non-existing word
	if trie.Delete("world") {
		t.Error("Expected to fail when deleting non-existing word 'world'")
	}

	// Delete word that is prefix of another
	if !trie.Delete("he") {
		t.Error("Expected to successfully delete 'he'")
	}

	if trie.Search("he") {
		t.Error("Expected 'he' to be deleted")
	}

	if !trie.Search("help") || !trie.Search("hero") || !trie.Search("her") {
		t.Error("Expected other words starting with 'he' to still exist")
	}
}

func TestTrieGetAllWords(t *testing.T) {
	trie := NewTrie()

	// Insert words in non-alphabetical order
	words := []string{"zebra", "apple", "banana", "app", "application"}
	for _, word := range words {
		trie.Insert(word)
	}

	allWords := trie.GetAllWords()

	// Expected words in lexicographical order
	expected := []string{"app", "apple", "application", "banana", "zebra"}

	if !reflect.DeepEqual(allWords, expected) {
		t.Errorf("Expected %v, got %v", expected, allWords)
	}
}

func TestTrieGetWordsWithPrefix(t *testing.T) {
	trie := NewTrie()

	// Insert words
	words := []string{"apple", "app", "application", "apply", "banana", "band", "bandana"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test prefix "app"
	appWords := trie.GetWordsWithPrefix("app")
	expectedApp := []string{"app", "apple", "application", "apply"}
	if !reflect.DeepEqual(appWords, expectedApp) {
		t.Errorf("Expected %v for prefix 'app', got %v", expectedApp, appWords)
	}

	// Test prefix "ban"
	banWords := trie.GetWordsWithPrefix("ban")
	expectedBan := []string{"banana", "band", "bandana"}
	if !reflect.DeepEqual(banWords, expectedBan) {
		t.Errorf("Expected %v for prefix 'ban', got %v", expectedBan, banWords)
	}

	// Test non-existing prefix
	nonExisting := trie.GetWordsWithPrefix("xyz")
	if len(nonExisting) != 0 {
		t.Errorf("Expected empty slice for non-existing prefix, got %v", nonExisting)
	}

	// Test empty prefix (should return all words)
	allWords := trie.GetWordsWithPrefix("")
	expectedAll := []string{"app", "apple", "application", "apply", "banana", "band", "bandana"}
	if !reflect.DeepEqual(allWords, expectedAll) {
		t.Errorf("Expected %v for empty prefix, got %v", expectedAll, allWords)
	}
}

func TestTrieClear(t *testing.T) {
	trie := NewTrie()

	// Insert words
	words := []string{"hello", "world", "test"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Verify words exist
	if trie.Len() != 3 {
		t.Errorf("Expected length 3 before clear, got %d", trie.Len())
	}

	// Clear the trie
	trie.Clear()

	// Verify trie is empty
	if trie.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", trie.Len())
	}

	for _, word := range words {
		if trie.Search(word) {
			t.Errorf("Expected word '%s' to be cleared", word)
		}
	}

	allWords := trie.GetAllWords()
	if len(allWords) != 0 {
		t.Errorf("Expected empty slice after clear, got %v", allWords)
	}
}

func TestTrieEmptyString(t *testing.T) {
	trie := NewTrie()

	// Test inserting empty string
	trie.Insert("")
	if trie.Len() != 0 {
		t.Errorf("Expected length 0 after inserting empty string, got %d", trie.Len())
	}

	// Test searching for empty string
	if trie.Search("") {
		t.Error("Expected false when searching for empty string")
	}

	// Test deleting empty string
	if trie.Delete("") {
		t.Error("Expected false when deleting empty string")
	}
}

func TestTrieDuplicateInsert(t *testing.T) {
	trie := NewTrie()

	// Insert same word multiple times
	trie.Insert("hello")
	trie.Insert("hello")
	trie.Insert("hello")

	// Length should still be 1
	if trie.Len() != 1 {
		t.Errorf("Expected length 1 after multiple inserts of same word, got %d", trie.Len())
	}

	if !trie.Search("hello") {
		t.Error("Expected to find 'hello' after multiple inserts")
	}
}

func TestTrieUnicodeSupport(t *testing.T) {
	trie := NewTrie()

	// Insert Unicode words
	words := []string{"café", "résumé", "naïve", "测试", "こんにちは"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test searching
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find Unicode word '%s'", word)
		}
	}

	// Test prefix search
	if !trie.StartsWith("caf") {
		t.Error("Expected to find prefix 'caf' for 'café'")
	}

	if !trie.StartsWith("测") {
		t.Error("Expected to find Unicode prefix '测' for '测试'")
	}
}

func TestTrieComplexScenario(t *testing.T) {
	trie := NewTrie()

	// Insert a complex set of related words
	words := []string{
		"a", "an", "and", "ant", "any", "app", "apple", "apply", "application",
		"be", "bee", "been", "beer", "best", "better",
		"cat", "car", "card", "care", "careful", "carefully",
	}

	for _, word := range words {
		trie.Insert(word)
	}

	// Test various operations
	if trie.Len() != len(words) {
		t.Errorf("Expected length %d, got %d", len(words), trie.Len())
	}

	// Test prefix searches
	prefixTests := []struct {
		prefix   string
		expected []string
	}{
		{"app", []string{"app", "apple", "application", "apply"}},
		{"be", []string{"be", "bee", "been", "beer", "best", "better"}},
		{"car", []string{"car", "card", "care", "careful", "carefully"}},
	}

	for _, test := range prefixTests {
		result := trie.GetWordsWithPrefix(test.prefix)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For prefix '%s', expected %v, got %v", test.prefix, test.expected, result)
		}
	}

	// Delete some words and verify structure integrity
	trie.Delete("application")
	trie.Delete("carefully")
	trie.Delete("bee")

	if trie.Len() != len(words)-3 {
		t.Errorf("Expected length %d after deletions, got %d", len(words)-3, trie.Len())
	}

	// Verify remaining words
	if !trie.Search("app") || !trie.Search("apply") {
		t.Error("Expected 'app' and 'apply' to remain after deleting 'application'")
	}

	if !trie.Search("care") || !trie.Search("careful") {
		t.Error("Expected 'care' and 'careful' to remain after deleting 'carefully'")
	}

	if !trie.Search("be") || !trie.Search("been") || !trie.Search("beer") {
		t.Error("Expected other 'be' words to remain after deleting 'bee'")
	}
}