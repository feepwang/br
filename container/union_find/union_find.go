//go:build !go1.23
// +build !go1.23

// Package union_find provides a Union-Find (Disjoint Set Union) data structure implementation.
// This file implements the Interface using union by rank and path compression optimizations.
package union_find

import (
	"sort"
)

// UnionFind implements the Interface using union by rank and path compression.
// It provides near-constant time operations for union and find.
type UnionFind struct {
	parent []int // parent[i] is the parent of element i
	rank   []int // rank[i] is the rank (approximate depth) of the tree rooted at i
	count  int   // number of disjoint sets
	size   int   // total number of elements
}

// NewUnionFind creates a new UnionFind with n elements.
// Initially, each element is in its own set.
func NewUnionFind(n int) *UnionFind {
	if n < 0 {
		n = 0
	}

	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
		size:   n,
	}

	// Initialize each element as its own parent
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}

	return uf
}

// Find returns the representative (root) of the set containing element x.
// Uses path compression for optimization.
func (uf *UnionFind) Find(x int) int {
	if x < 0 || x >= uf.size {
		return -1 // Invalid element
	}

	// Path compression: make every node point directly to the root
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

// Union merges the sets containing elements x and y.
// Uses union by rank for optimization.
func (uf *UnionFind) Union(x, y int) {
	if x < 0 || x >= uf.size || y < 0 || y >= uf.size {
		return // Invalid elements
	}

	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// Already in the same set
	if rootX == rootY {
		return
	}

	// Union by rank: attach smaller tree under root of larger tree
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.count--
}

// Connected returns true if elements x and y belong to the same set.
func (uf *UnionFind) Connected(x, y int) bool {
	if x < 0 || x >= uf.size || y < 0 || y >= uf.size {
		return false
	}
	return uf.Find(x) == uf.Find(y)
}

// Count returns the number of disjoint sets.
func (uf *UnionFind) Count() int {
	return uf.count
}

// Size returns the total number of elements.
func (uf *UnionFind) Size() int {
	return uf.size
}

// SetSize returns the size of the set containing element x.
func (uf *UnionFind) SetSize(x int) int {
	if x < 0 || x >= uf.size {
		return 0
	}

	root := uf.Find(x)
	setSize := 0

	for i := 0; i < uf.size; i++ {
		if uf.Find(i) == root {
			setSize++
		}
	}

	return setSize
}

// Sets returns all disjoint sets as a slice of slices.
// Each inner slice contains elements belonging to the same set.
func (uf *UnionFind) Sets() [][]int {
	// Group elements by their root
	groups := make(map[int][]int)

	for i := 0; i < uf.size; i++ {
		root := uf.Find(i)
		groups[root] = append(groups[root], i)
	}

	// Convert map to slice of slices
	result := make([][]int, 0, len(groups))
	roots := make([]int, 0, len(groups))

	// Sort roots for consistent output
	for root := range groups {
		roots = append(roots, root)
	}
	sort.Ints(roots)

	for _, root := range roots {
		set := groups[root]
		sort.Ints(set) // Sort elements within each set
		result = append(result, set)
	}

	return result
}

// Reset reinitializes the Union-Find structure with the given size.
// All elements will be in separate sets after reset.
func (uf *UnionFind) Reset(n int) {
	if n < 0 {
		n = 0
	}

	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	uf.count = n
	uf.size = n

	// Initialize each element as its own parent
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
}
