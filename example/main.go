package main

import (
	"fmt"

	"github.com/feepwang/br/container/ordered_map"
	"github.com/feepwang/br/container/pair"
	"github.com/feepwang/br/container/trie_tree"
)

func main() {
	fmt.Println("Brain Rehabilitation Toolkit Demo")
	fmt.Println("=================================")

	// Demonstrate Pair usage
	fmt.Println("\n1. Pair Demo:")
	p := pair.Pair[string, int]{First: "score", Second: 95}
	fmt.Printf("   Pair: %s = %d\n", p.First, p.Second)

	// Demonstrate OrderedMap usage
	fmt.Println("\n2. OrderedMap (Red-Black Tree) Demo:")
	
	tree := ordered_map.NewRedBlackTree[int, string]()
	
	// Insert some values
	fmt.Println("   Inserting values...")
	tree.Set(5, "five")
	tree.Set(2, "two")
	tree.Set(8, "eight")
	tree.Set(1, "one")
	tree.Set(7, "seven")
	tree.Set(3, "three")
	
	fmt.Printf("   Tree size: %d\n", tree.Len())
	fmt.Printf("   Tree capacity: %d\n", tree.Cap())
	
	// Demonstrate retrieval
	if val, exists := tree.Get(5); exists {
		fmt.Printf("   Key 5 -> %s\n", val)
	}
	
	// Demonstrate ordering
	fmt.Println("   Keys in order:", tree.Keys())
	fmt.Println("   Values in order:", tree.Values())
	
	// Demonstrate deletion
	fmt.Println("   Deleting key 2...")
	tree.Delete(2)
	fmt.Println("   Keys after deletion:", tree.Keys())
	
	// Demonstrate pairs
	fmt.Println("   Key-Value pairs:")
	for _, p := range tree.Pairs() {
		fmt.Printf("     %d -> %s\n", p.First, p.Second)
	}
	
	// Demonstrate Trie Tree usage
	fmt.Println("\n3. Trie Tree Demo:")
	
	trie := trie_tree.NewTrie()
	
	// Insert some words
	words := []string{"apple", "app", "application", "apply", "banana", "band", "bandana"}
	fmt.Println("   Inserting words...")
	for _, word := range words {
		trie.Insert(word)
	}
	
	fmt.Printf("   Trie size: %d\n", trie.Len())
	
	// Demonstrate search
	searchWords := []string{"app", "apple", "apply", "orange"}
	fmt.Println("   Search results:")
	for _, word := range searchWords {
		exists := trie.Search(word)
		fmt.Printf("     '%s' exists: %t\n", word, exists)
	}
	
	// Demonstrate prefix search
	prefixes := []string{"app", "ban", "xyz"}
	fmt.Println("   Prefix search results:")
	for _, prefix := range prefixes {
		hasPrefix := trie.StartsWith(prefix)
		fmt.Printf("     Words starting with '%s': %t\n", prefix, hasPrefix)
	}
	
	// Demonstrate getting words with prefix
	fmt.Println("   Words starting with 'app':")
	appWords := trie.GetWordsWithPrefix("app")
	for _, word := range appWords {
		fmt.Printf("     %s\n", word)
	}
	
	// Demonstrate deletion
	fmt.Println("   Deleting 'application'...")
	deleted := trie.Delete("application")
	fmt.Printf("   Deletion successful: %t\n", deleted)
	fmt.Printf("   Trie size after deletion: %d\n", trie.Len())
	
	// Show all remaining words
	fmt.Println("   All remaining words:")
	allWords := trie.GetAllWords()
	for _, word := range allWords {
		fmt.Printf("     %s\n", word)
	}
	
	fmt.Println("\nDemo completed successfully!")
}