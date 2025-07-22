//go:build !go1.23
// +build !go1.23

// Package union_find provides a Union-Find (Disjoint Set Union) data structure implementation.
// Union-Find is a data structure that tracks a set of elements partitioned into
// non-overlapping disjoint sets. It provides near-constant time operations for
// union and find operations using path compression and union by rank optimizations.
package union_find

// Interface defines the operations for a Union-Find data structure.
// Union-Find efficiently maintains disjoint sets and supports fast union and find operations.
type Interface interface {
	// Union merges the sets containing elements x and y.
	// If x and y are already in the same set, this operation has no effect.
	Union(x, y int)

	// Find returns the representative (root) of the set containing element x.
	// Uses path compression for optimization.
	Find(x int) int

	// Connected returns true if elements x and y belong to the same set.
	Connected(x, y int) bool

	// Count returns the number of disjoint sets.
	Count() int

	// Size returns the total number of elements.
	Size() int

	// SetSize returns the size of the set containing element x.
	SetSize(x int) int

	// Sets returns all disjoint sets as a slice of slices.
	// Each inner slice contains elements belonging to the same set.
	Sets() [][]int

	// Reset reinitializes the Union-Find structure with the given size.
	// All elements will be in separate sets after reset.
	Reset(size int)
}
