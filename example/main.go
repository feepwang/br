package main

import (
	"fmt"

	"github.com/feepwang/br/container/ordered_map"
	"github.com/feepwang/br/container/pair"
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
	
	fmt.Println("\nDemo completed successfully!")
}