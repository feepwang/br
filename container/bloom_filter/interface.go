//go:build !go1.23
// +build !go1.23

// Package bloom_filter provides a Bloom Filter data structure implementation.
// A Bloom Filter is a probabilistic data structure that tests whether an element
// is a member of a set. False positive matches are possible, but false negatives
// are not. It's space-efficient and provides constant-time operations.
package bloom_filter

// Interface defines the operations for a Bloom Filter data structure.
// A Bloom Filter uses a bit array and multiple hash functions to efficiently
// test set membership with a configurable false positive rate.
type Interface[T comparable] interface {
	// Add inserts an item into the Bloom filter.
	// This operation never fails and always succeeds.
	// Time complexity: O(k) where k is the number of hash functions.
	Add(item T)

	// Contains tests whether an item might be in the set.
	// Returns true if the item might be in the set (possible false positive).
	// Returns false if the item is definitely not in the set (no false negatives).
	// Time complexity: O(k) where k is the number of hash functions.
	Contains(item T) bool

	// Clear resets the Bloom filter to its initial empty state.
	// All previously added items will be forgotten.
	// Time complexity: O(m) where m is the size of the bit array.
	Clear()

	// Len returns the approximate number of items that have been added to the filter.
	// This is an estimation based on the number of set bits and may not be exact
	// due to hash collisions and the probabilistic nature of Bloom filters.
	// Time complexity: O(1).
	Len() int

	// Capacity returns the estimated maximum number of items that can be added
	// before the false positive rate exceeds the configured threshold.
	// Time complexity: O(1).
	Capacity() int

	// FalsePositiveRate returns the current estimated false positive rate
	// based on the number of items added and the filter configuration.
	// Time complexity: O(1).
	FalsePositiveRate() float64

	// BitSize returns the size of the underlying bit array.
	// Time complexity: O(1).
	BitSize() int

	// HashCount returns the number of hash functions used by the filter.
	// Time complexity: O(1).
	HashCount() int
}
