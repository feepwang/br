//go:build go1.23
// +build go1.23

// Package ordered_map provides go1.23-specific methods for RedBlackTree.
// This file adds iter.Seq related methods for Interface.

package ordered_map

import (
	"cmp"
	"iter"
)

// KeySeq returns an iterator for keys (go1.23).
// Uses efficient iterative in-order traversal without pre-allocating all keys.
func (t *RedBlackTree[K, V]) KeySeq() iter.Seq[K] {
	return func(yield func(K) bool) {
		inOrderKeysIterative(t.root, yield)
	}
}

// ValueSeq returns an iterator for values (go1.23).
// Uses efficient iterative in-order traversal without pre-allocating all values.
func (t *RedBlackTree[K, V]) ValueSeq() iter.Seq[V] {
	return func(yield func(V) bool) {
		inOrderValuesIterative(t.root, yield)
	}
}

// PairSeq returns an iterator for key-value pairs (go1.23).
// Uses efficient iterative in-order traversal without pre-allocating all pairs.
func (t *RedBlackTree[K, V]) PairSeq() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		inOrderPairsIterative(t.root, yield)
	}
}

// inOrderKeysIterative performs iterative in-order traversal for keys.
func inOrderKeysIterative[K cmp.Ordered, V any](root *rbNode[K, V], yield func(K) bool) {
	if root == nil {
		return
	}

	stack := make([]*rbNode[K, V], 0)
	current := root

	for len(stack) > 0 || current != nil {
		// Go to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		// Current must be nil at this point
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Yield the key
		if !yield(current.key) {
			return
		}

		// Visit the right subtree
		current = current.right
	}
}

// inOrderValuesIterative performs iterative in-order traversal for values.
func inOrderValuesIterative[K cmp.Ordered, V any](root *rbNode[K, V], yield func(V) bool) {
	if root == nil {
		return
	}

	stack := make([]*rbNode[K, V], 0)
	current := root

	for len(stack) > 0 || current != nil {
		// Go to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		// Current must be nil at this point
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Yield the value
		if !yield(current.value) {
			return
		}

		// Visit the right subtree
		current = current.right
	}
}

// inOrderPairsIterative performs iterative in-order traversal for key-value pairs.
func inOrderPairsIterative[K cmp.Ordered, V any](root *rbNode[K, V], yield func(K, V) bool) {
	if root == nil {
		return
	}

	stack := make([]*rbNode[K, V], 0)
	current := root

	for len(stack) > 0 || current != nil {
		// Go to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		// Current must be nil at this point
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Yield the key-value pair
		if !yield(current.key, current.value) {
			return
		}

		// Visit the right subtree
		current = current.right
	}
}
