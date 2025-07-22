//go:build go1.23
// +build go1.23

// Package union_find provides a Union-Find (Disjoint Set Union) data structure implementation.
// This version includes Go 1.23+ features such as iterator support.
package union_find

import (
	"iter"
)

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

	// AllSets returns an iterator over all disjoint sets.
	// Each iteration yields a slice containing elements of one set.
	AllSets() iter.Seq[[]int]

	// SetMembers returns an iterator over all members of the set containing element x.
	SetMembers(x int) iter.Seq[int]
}
