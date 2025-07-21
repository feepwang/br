//go:build go1.23
// +build go1.23

// Package set provides a generic set data structure implementation.
// A set is a collection of unique elements that supports mathematical
// set operations like union, intersection, and difference.
package set

import "iter"

// Interface defines the operations for a generic set data structure.
// A set maintains a collection of unique elements and supports efficient
// membership testing and set operations.
type Interface[T comparable] interface {
	// Add inserts an element into the set.
	// Returns true if the element was newly added, false if it already existed.
	Add(element T) bool

	// Remove deletes an element from the set.
	// Returns true if the element was found and removed, false if it didn't exist.
	Remove(element T) bool

	// Contains checks if an element exists in the set.
	Contains(element T) bool

	// Len returns the number of elements in the set.
	Len() int

	// IsEmpty returns true if the set contains no elements.
	IsEmpty() bool

	// Clear removes all elements from the set.
	Clear()

	// Slice returns all elements as a slice in no particular order.
	Slice() []T

	// Equal returns true if this set contains exactly the same elements as other.
	Equal(other Interface[T]) bool

	// IsSubset returns true if all elements in this set are contained in other.
	IsSubset(other Interface[T]) bool

	// IsSuperset returns true if this set contains all elements from other.
	IsSuperset(other Interface[T]) bool

	// Union returns a new set containing all elements from both sets.
	Union(other Interface[T]) Interface[T]

	// Intersection returns a new set containing elements present in both sets.
	Intersection(other Interface[T]) Interface[T]

	// Difference returns a new set containing elements in this set but not in other.
	Difference(other Interface[T]) Interface[T]

	// SymmetricDifference returns a new set containing elements in either set but not in both.
	SymmetricDifference(other Interface[T]) Interface[T]

	// All returns an iterator over all elements in the set.
	All() iter.Seq[T]
}