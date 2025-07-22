//go:build go1.23
// +build go1.23

package skip_list

import (
	"cmp"
	"iter"
	"math/rand"
	"time"

	"github.com/feepwang/br/container/pair"
)

const (
	// maxLevel defines the maximum number of levels in the skip list.
	// This limits the height to prevent excessive memory usage.
	maxLevel = 32

	// probability defines the probability of a node having a pointer at the next level.
	// Traditional skip lists use p = 0.5, which provides good balance between
	// search time and space usage.
	probability = 0.5
)

// node represents a single node in the skip list.
type node[K comparable, V any] struct {
	key     K
	value   V
	forward []*node[K, V] // Array of forward pointers for each level
}

// SkipList is a concrete implementation of the Interface.
type SkipList[K comparable, V any] struct {
	header  *node[K, V]      // Header node (sentinel)
	level   int              // Current maximum level of the list
	length  int              // Number of elements in the list
	rng     *rand.Rand       // Random number generator for level assignment
	compare func(a, b K) int // Comparison function for keys
}

// NewSkipList creates and returns a new empty skip list.
func NewSkipList[K comparable, V any](compare func(a, b K) int) Interface[K, V] {
	header := &node[K, V]{
		forward: make([]*node[K, V], maxLevel),
	}

	return &SkipList[K, V]{
		header:  header,
		level:   0,
		length:  0,
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
		compare: compare,
	}
}

// NewOrderedSkipList creates a new skip list for ordered types (types that implement cmp.Ordered).
func NewOrderedSkipList[K cmp.Ordered, V any]() Interface[K, V] {
	return NewSkipList[K, V](cmp.Compare[K])
}

// randomLevel generates a random level for a new node.
// Uses geometric distribution with the specified probability.
func (sl *SkipList[K, V]) randomLevel() int {
	level := 0
	for sl.rng.Float64() < probability && level < maxLevel-1 {
		level++
	}
	return level
}

// search finds the position where a key should be inserted or already exists.
// Returns the update array needed for insertion/deletion operations.
func (sl *SkipList[K, V]) search(key K) ([]*node[K, V], *node[K, V]) {
	update := make([]*node[K, V], maxLevel)
	current := sl.header

	// Start from the highest level and work downward
	for i := sl.level; i >= 0; i-- {
		// Move forward while the next node's key is less than the search key
		for current.forward[i] != nil && sl.compare(current.forward[i].key, key) < 0 {
			current = current.forward[i]
		}
		update[i] = current
	}

	// Move to the next node (potential match)
	current = current.forward[0]
	return update, current
}

// Len returns the number of key-value pairs stored in the skip list.
func (sl *SkipList[K, V]) Len() int {
	return sl.length
}

// Get retrieves the value associated with the given key.
func (sl *SkipList[K, V]) Get(key K) (V, bool) {
	_, current := sl.search(key)
	if current != nil && sl.compare(current.key, key) == 0 {
		return current.value, true
	}
	var zero V
	return zero, false
}

// GetMutable returns a pointer to the value associated with the given key.
func (sl *SkipList[K, V]) GetMutable(key K) (*V, bool) {
	_, current := sl.search(key)
	if current != nil && sl.compare(current.key, key) == 0 {
		return &current.value, true
	}
	return nil, false
}

// Set inserts or updates a key-value pair in the skip list.
func (sl *SkipList[K, V]) Set(key K, value V) {
	update, current := sl.search(key)

	// If key already exists, update the value
	if current != nil && sl.compare(current.key, key) == 0 {
		current.value = value
		return
	}

	// Generate random level for the new node
	newLevel := sl.randomLevel()

	// If new level is higher than current level, update the header pointers
	if newLevel > sl.level {
		for i := sl.level + 1; i <= newLevel; i++ {
			update[i] = sl.header
		}
		sl.level = newLevel
	}

	// Create new node
	newNode := &node[K, V]{
		key:     key,
		value:   value,
		forward: make([]*node[K, V], newLevel+1),
	}

	// Update forward pointers
	for i := 0; i <= newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}

	sl.length++
}

// Delete removes the key-value pair with the given key from the skip list.
func (sl *SkipList[K, V]) Delete(key K) bool {
	update, current := sl.search(key)

	// If key doesn't exist, return false
	if current == nil || sl.compare(current.key, key) != 0 {
		return false
	}

	// Update forward pointers to skip the node being deleted
	for i := 0; i <= sl.level; i++ {
		if update[i].forward[i] != current {
			break
		}
		update[i].forward[i] = current.forward[i]
	}

	// Update the level of the skip list if necessary
	for sl.level > 0 && sl.header.forward[sl.level] == nil {
		sl.level--
	}

	sl.length--
	return true
}

// Has checks whether the given key exists in the skip list.
func (sl *SkipList[K, V]) Has(key K) bool {
	_, exists := sl.Get(key)
	return exists
}

// Clear removes all key-value pairs from the skip list.
func (sl *SkipList[K, V]) Clear() {
	sl.header.forward = make([]*node[K, V], maxLevel)
	sl.level = 0
	sl.length = 0
}

// Keys returns a slice of all keys in the skip list in sorted order.
func (sl *SkipList[K, V]) Keys() []K {
	keys := make([]K, 0, sl.length)
	current := sl.header.forward[0]
	for current != nil {
		keys = append(keys, current.key)
		current = current.forward[0]
	}
	return keys
}

// Values returns a slice of all values in the skip list in the order of their keys.
func (sl *SkipList[K, V]) Values() []V {
	values := make([]V, 0, sl.length)
	current := sl.header.forward[0]
	for current != nil {
		values = append(values, current.value)
		current = current.forward[0]
	}
	return values
}

// Pairs returns a slice of all key-value pairs in the skip list in sorted order by key.
func (sl *SkipList[K, V]) Pairs() []pair.Pair[K, V] {
	pairs := make([]pair.Pair[K, V], 0, sl.length)
	current := sl.header.forward[0]
	for current != nil {
		pairs = append(pairs, pair.Pair[K, V]{First: current.key, Second: current.value})
		current = current.forward[0]
	}
	return pairs
}

// Range calls the provided function for each key-value pair in sorted order by key.
func (sl *SkipList[K, V]) Range(fn func(key K, value V) bool) {
	current := sl.header.forward[0]
	for current != nil {
		if !fn(current.key, current.value) {
			break
		}
		current = current.forward[0]
	}
}

// RangeFrom calls the provided function for key-value pairs starting from the given key.
func (sl *SkipList[K, V]) RangeFrom(start K, fn func(key K, value V) bool) {
	// Find the first node with key >= start
	current := sl.header
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && sl.compare(current.forward[i].key, start) < 0 {
			current = current.forward[i]
		}
	}
	current = current.forward[0]

	// Iterate from the starting position
	for current != nil {
		if !fn(current.key, current.value) {
			break
		}
		current = current.forward[0]
	}
}

// RangeBetween calls the provided function for key-value pairs within the given range.
func (sl *SkipList[K, V]) RangeBetween(start, end K, fn func(key K, value V) bool) {
	// Determine the logical start and end based on comparator
	actualStart, actualEnd := start, end
	if sl.compare(start, end) > 0 {
		actualStart, actualEnd = end, start
	}

	// Find the first node with key >= actualStart
	current := sl.header
	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && sl.compare(current.forward[i].key, actualStart) < 0 {
			current = current.forward[i]
		}
	}
	current = current.forward[0]

	// Iterate while key <= actualEnd
	for current != nil && sl.compare(current.key, actualEnd) <= 0 {
		if !fn(current.key, current.value) {
			break
		}
		current = current.forward[0]
	}
}

// All returns an iterator over all key-value pairs in sorted order by key.
func (sl *SkipList[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		current := sl.header.forward[0]
		for current != nil {
			if !yield(current.key, current.value) {
				return
			}
			current = current.forward[0]
		}
	}
}

// AllFrom returns an iterator over key-value pairs starting from the given key.
func (sl *SkipList[K, V]) AllFrom(start K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		// Find the first node with key >= start
		current := sl.header
		for i := sl.level; i >= 0; i-- {
			for current.forward[i] != nil && sl.compare(current.forward[i].key, start) < 0 {
				current = current.forward[i]
			}
		}
		current = current.forward[0]

		// Iterate from the starting position
		for current != nil {
			if !yield(current.key, current.value) {
				return
			}
			current = current.forward[0]
		}
	}
}

// AllBetween returns an iterator over key-value pairs within the given range.
func (sl *SkipList[K, V]) AllBetween(start, end K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		// Determine the logical start and end based on comparator
		// If start > end according to the comparator, swap them
		actualStart, actualEnd := start, end
		if sl.compare(start, end) > 0 {
			actualStart, actualEnd = end, start
		}

		// Find the first node with key >= actualStart
		current := sl.header
		for i := sl.level; i >= 0; i-- {
			for current.forward[i] != nil && sl.compare(current.forward[i].key, actualStart) < 0 {
				current = current.forward[i]
			}
		}
		current = current.forward[0]

		// Iterate while key <= actualEnd
		for current != nil && sl.compare(current.key, actualEnd) <= 0 {
			if !yield(current.key, current.value) {
				return
			}
			current = current.forward[0]
		}
	}
}
