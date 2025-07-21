//go:build !go1.23
// +build !go1.23

// Package skip_list provides a Skip List data structure implementation.
// A Skip List is a probabilistic data structure that allows for efficient
// search, insertion, and deletion operations with average O(log n) time complexity.
// It maintains elements in sorted order and uses multiple levels for fast traversal.
package skip_list

import (
	"cmp"

	"github.com/feepwang/br/container/pair"
)

// Interface defines the operations for a Skip List data structure.
// A Skip List maintains key-value pairs in sorted order by key and provides
// efficient operations through a probabilistic multi-level structure.
type Interface[K cmp.Ordered, V any] interface {
	// Len returns the number of key-value pairs stored in the skip list.
	Len() int

	// Get retrieves the value associated with the given key.
	// Returns the value and true if the key exists, zero value and false otherwise.
	Get(key K) (V, bool)

	// GetMutable returns a pointer to the value associated with the given key.
	// This allows for in-place modification of the value.
	// Returns a pointer to the value and true if the key exists, nil and false otherwise.
	GetMutable(key K) (*V, bool)

	// Set inserts or updates a key-value pair in the skip list.
	// If the key already exists, its value is updated.
	Set(key K, value V)

	// Delete removes the key-value pair with the given key from the skip list.
	// Returns true if the key was found and removed, false otherwise.
	Delete(key K) bool

	// Has checks whether the given key exists in the skip list.
	Has(key K) bool

	// Clear removes all key-value pairs from the skip list.
	Clear()

	// Keys returns a slice of all keys in the skip list in sorted order.
	Keys() []K

	// Values returns a slice of all values in the skip list in the order of their keys.
	Values() []V

	// Pairs returns a slice of all key-value pairs in the skip list in sorted order by key.
	Pairs() []pair.Pair[K, V]

	// Range calls the provided function for each key-value pair in the skip list
	// in sorted order by key. If the function returns false, the iteration stops.
	Range(fn func(key K, value V) bool)

	// RangeFrom calls the provided function for each key-value pair in the skip list
	// starting from the given key (inclusive) in sorted order by key.
	// If the function returns false, the iteration stops.
	RangeFrom(start K, fn func(key K, value V) bool)

	// RangeBetween calls the provided function for each key-value pair in the skip list
	// within the given key range [start, end] (both inclusive) in sorted order by key.
	// If the function returns false, the iteration stops.
	RangeBetween(start, end K, fn func(key K, value V) bool)
}