//go:build go1.23
// +build go1.23

// Package ordered_map provides go1.23-specific methods for RedBlackTree.
// This file adds iter.Seq related methods for Interface.

package ordered_map

import (
	"iter"
)

// KeySeq returns an iterator for keys (go1.23).
func (t *RedBlackTree[K, V]) KeySeq() iter.Seq[K] {
	keys := t.Keys()
	return func(yield func(K) bool) {
		for _, k := range keys {
			if !yield(k) {
				return
			}
		}
	}
}

// ValueSeq returns an iterator for values (go1.23).
func (t *RedBlackTree[K, V]) ValueSeq() iter.Seq[V] {
	values := t.Values()
	return func(yield func(V) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}

// PairSeq returns an iterator for key-value pairs (go1.23).
func (t *RedBlackTree[K, V]) PairSeq() iter.Seq2[K, V] {
	pairs := t.Pairs()
	return func(yield func(K, V) bool) {
		for _, p := range pairs {
			if !yield(p.First, p.Second) {
				return
			}
		}
	}
}
