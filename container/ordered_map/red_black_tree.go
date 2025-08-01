// Package ordered_map provides an ordered map implementation using Red-Black Tree.
// This file implements the Interface[K, V] using a Red-Black Tree.
// Comments are added at key places for beginners.

package ordered_map

import (
	"cmp"

	"github.com/feepwang/br/container/pair"
)

// Color represents the color of a Red-Black Tree node.
type color bool

const (
	red   color = true
	black color = false
)

// rbNode is a node in the Red-Black Tree.
type rbNode[K cmp.Ordered, V any] struct {
	key    K
	value  V
	left   *rbNode[K, V]
	right  *rbNode[K, V]
	parent *rbNode[K, V]
	color  color
}

// RedBlackTree implements the ordered_map.Interface using a Red-Black Tree.
type RedBlackTree[K cmp.Ordered, V any] struct {
	root *rbNode[K, V]
	size int
}

// NewRedBlackTree creates a new RedBlackTree.
func NewRedBlackTree[K cmp.Ordered, V any]() *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{}
}

// Len returns the number of elements in the map.
func (t *RedBlackTree[K, V]) Len() int {
	return t.size
}

// Cap returns the capacity of the map. For Red-Black Tree, capacity equals size since it's dynamic.
func (t *RedBlackTree[K, V]) Cap() int {
	return t.size
}

// Get searches for a key and returns its value and existence.
func (t *RedBlackTree[K, V]) Get(key K) (V, bool) {
	n := t.root
	for n != nil {
		if cmp.Less(key, n.key) {
			n = n.left
		} else if cmp.Less(n.key, key) {
			n = n.right
		} else {
			return n.value, true
		}
	}
	var zero V
	return zero, false
}

// GetMutable returns a pointer to the value for mutation.
func (t *RedBlackTree[K, V]) GetMutable(key K) (*V, bool) {
	n := t.root
	for n != nil {
		if cmp.Less(key, n.key) {
			n = n.left
		} else if cmp.Less(n.key, key) {
			n = n.right
		} else {
			return &n.value, true
		}
	}
	return nil, false
}

// Set inserts or updates a key-value pair.
func (t *RedBlackTree[K, V]) Set(key K, value V) {
	// Standard BST insert, then fixup for Red-Black properties
	var inserted *rbNode[K, V]
	if t.root == nil {
		// Tree is empty, insert root
		inserted = &rbNode[K, V]{key: key, value: value, color: black}
		t.root = inserted
		t.size++
		return
	}
	n := t.root
	var parent *rbNode[K, V]
	for n != nil {
		parent = n
		if cmp.Less(key, n.key) {
			n = n.left
		} else if cmp.Less(n.key, key) {
			n = n.right
		} else {
			// Key exists, update value
			n.value = value
			return
		}
	}
	inserted = &rbNode[K, V]{key: key, value: value, parent: parent, color: red}
	if cmp.Less(key, parent.key) {
		parent.left = inserted
	} else {
		parent.right = inserted
	}
	t.size++
	// Fix Red-Black Tree properties after insert
	fixInsert(t, inserted)
}

// fixInsert restores Red-Black Tree properties after insertion.
func fixInsert[K cmp.Ordered, V any](t *RedBlackTree[K, V], n *rbNode[K, V]) {
	// Key place: Red-Black Tree balancing after insert
	for n != t.root && n.parent.color == red {
		if n.parent == n.parent.parent.left {
			uncle := n.parent.parent.right
			if uncle != nil && uncle.color == red {
				n.parent.color = black
				uncle.color = black
				n.parent.parent.color = red
				n = n.parent.parent
			} else {
				if n == n.parent.right {
					n = n.parent
					rotateLeft(t, n)
				}
				n.parent.color = black
				n.parent.parent.color = red
				rotateRight(t, n.parent.parent)
			}
		} else {
			uncle := n.parent.parent.left
			if uncle != nil && uncle.color == red {
				n.parent.color = black
				uncle.color = black
				n.parent.parent.color = red
				n = n.parent.parent
			} else {
				if n == n.parent.left {
					n = n.parent
					rotateRight(t, n)
				}
				n.parent.color = black
				n.parent.parent.color = red
				rotateLeft(t, n.parent.parent)
			}
		}
	}
	t.root.color = black
}

// rotateLeft performs a left rotation.
func rotateLeft[K cmp.Ordered, V any](t *RedBlackTree[K, V], x *rbNode[K, V]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// rotateRight performs a right rotation.
func rotateRight[K cmp.Ordered, V any](t *RedBlackTree[K, V], x *rbNode[K, V]) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// Has checks if a key exists in the map.
func (t *RedBlackTree[K, V]) Has(key K) bool {
	_, ok := t.Get(key)
	return ok
}

// Delete removes a key from the map.
func (t *RedBlackTree[K, V]) Delete(key K) bool {
	// Key place: Red-Black Tree delete and fixup
	n := t.root
	for n != nil {
		if cmp.Less(key, n.key) {
			n = n.left
		} else if cmp.Less(n.key, key) {
			n = n.right
		} else {
			deleteNode(t, n)
			t.size--
			return true
		}
	}
	return false
}

// deleteNode removes a node and fixes Red-Black properties.
func deleteNode[K cmp.Ordered, V any](t *RedBlackTree[K, V], z *rbNode[K, V]) {
	// Standard BST delete, then fixup for Red-Black properties
	// Key place: For beginners, see Red-Black Tree delete algorithm for details.

	var y, x *rbNode[K, V]

	// Find the node to actually delete (y) and its replacement (x)
	if z.left == nil || z.right == nil {
		y = z // z has at most one child
	} else {
		// z has two children, find successor
		y = z.right
		for y.left != nil {
			y = y.left
		}
	}

	// Set x to y's child (or nil if y has no children)
	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	// Link x to y's parent
	if x != nil {
		x.parent = y.parent
	}

	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	// If y is not the node to delete, copy y's data to z
	if y != z {
		z.key = y.key
		z.value = y.value
	}

	// Fix Red-Black properties if a black node was deleted
	if y.color == black && x != nil {
		fixDelete(t, x)
	}
}

// fixDelete restores Red-Black Tree properties after deletion.
func fixDelete[K cmp.Ordered, V any](t *RedBlackTree[K, V], x *rbNode[K, V]) {
	for x != t.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right // sibling
			if w.color == red {
				w.color = black
				x.parent.color = red
				rotateLeft(t, x.parent)
				w = x.parent.right
			}
			if (w.left == nil || w.left.color == black) &&
				(w.right == nil || w.right.color == black) {
				w.color = red
				x = x.parent
			} else {
				if w.right == nil || w.right.color == black {
					if w.left != nil {
						w.left.color = black
					}
					w.color = red
					rotateRight(t, w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				if w.right != nil {
					w.right.color = black
				}
				rotateLeft(t, x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left // sibling
			if w.color == red {
				w.color = black
				x.parent.color = red
				rotateRight(t, x.parent)
				w = x.parent.left
			}
			if (w.right == nil || w.right.color == black) &&
				(w.left == nil || w.left.color == black) {
				w.color = red
				x = x.parent
			} else {
				if w.left == nil || w.left.color == black {
					if w.right != nil {
						w.right.color = black
					}
					w.color = red
					rotateLeft(t, w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				if w.left != nil {
					w.left.color = black
				}
				rotateRight(t, x.parent)
				x = t.root
			}
		}
	}
	x.color = black
}

// Keys returns all keys in order.
func (t *RedBlackTree[K, V]) Keys() []K {
	var keys []K
	inOrderKeys(t.root, &keys)
	return keys
}

func inOrderKeys[K cmp.Ordered, V any](n *rbNode[K, V], keys *[]K) {
	if n == nil {
		return
	}
	inOrderKeys(n.left, keys)
	*keys = append(*keys, n.key)
	inOrderKeys(n.right, keys)
}

// Values returns all values in order.
func (t *RedBlackTree[K, V]) Values() []V {
	var values []V
	inOrderValues(t.root, &values)
	return values
}

func inOrderValues[K cmp.Ordered, V any](n *rbNode[K, V], values *[]V) {
	if n == nil {
		return
	}
	inOrderValues(n.left, values)
	*values = append(*values, n.value)
	inOrderValues(n.right, values)
}

// Pairs returns all key-value pairs in order.
func (t *RedBlackTree[K, V]) Pairs() []pair.Pair[K, V] {
	var pairs []pair.Pair[K, V]
	inOrderPairs(t.root, &pairs)
	return pairs
}

func inOrderPairs[K cmp.Ordered, V any](n *rbNode[K, V], pairs *[]pair.Pair[K, V]) {
	if n == nil {
		return
	}
	inOrderPairs(n.left, pairs)
	*pairs = append(*pairs, pair.Pair[K, V]{First: n.key, Second: n.value})
	inOrderPairs(n.right, pairs)
}

// Ensure RedBlackTree implements Interface (for non-go1.23 version)
var _ Interface[int, int] = (*RedBlackTree[int, int])(nil)
