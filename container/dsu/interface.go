//go:build !go1.23
// +build !go1.23

// Package dsu provides a Disjoint Set Union (Union-Find) data structure implementation.
// A Disjoint Set Union maintains a collection of disjoint sets and supports efficient
// find and union operations with path compression and union by rank optimizations.
// It's commonly used for cycle detection, connectivity queries, and Kruskal's algorithm.
package dsu

// Interface defines the operations for a Disjoint Set Union data structure.
// A DSU maintains a collection of disjoint sets of integers from 0 to n-1
// and provides efficient operations to find set representatives and union sets.
type Interface interface {
	// Find returns the representative (root) of the set containing element x.
	// Uses path compression optimization to flatten the tree structure.
	// Time complexity: O(α(n)) amortized, where α is the inverse Ackermann function.
	Find(x int) int

	// Union merges the sets containing elements x and y.
	// Returns true if the elements were in different sets (union performed),
	// false if they were already in the same set.
	// Uses union by rank optimization to keep trees balanced.
	// Time complexity: O(α(n)) amortized.
	Union(x, y int) bool

	// Connected returns true if elements x and y are in the same set.
	// This is equivalent to Find(x) == Find(y) but more expressive.
	// Time complexity: O(α(n)) amortized.
	Connected(x, y int) bool

	// ComponentCount returns the number of disjoint sets (connected components).
	// Initially equals n, decreases by 1 with each successful union operation.
	// Time complexity: O(1).
	ComponentCount() int

	// Size returns the total number of elements in the DSU.
	// This is the n value used during initialization.
	// Time complexity: O(1).
	Size() int
}
