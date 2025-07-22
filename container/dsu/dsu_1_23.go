//go:build go1.23
// +build go1.23

package dsu

// DSU represents a Disjoint Set Union (Union-Find) data structure.
// It maintains a forest of trees where each tree represents a disjoint set.
// The structure uses path compression and union by rank optimizations
// to achieve nearly constant time complexity for operations.
type DSU struct {
	parent     []int // parent[i] is the parent of element i in the tree
	rank       []int // rank[i] is the approximate depth of the tree rooted at i
	components int   // number of disjoint components
	size       int   // total number of elements
}

// NewDSU creates a new Disjoint Set Union with n elements (0 to n-1).
// Initially, each element forms its own singleton set.
// Returns nil if n <= 0.
func NewDSU(n int) Interface {
	if n <= 0 {
		return nil
	}

	dsu := &DSU{
		parent:     make([]int, n),
		rank:       make([]int, n),
		components: n,
		size:       n,
	}

	// Initialize each element as its own parent (singleton sets)
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		// rank[i] = 0 (default zero value)
	}

	return dsu
}

// Find returns the representative (root) of the set containing element x.
// Implements path compression optimization: during traversal to the root,
// all nodes on the path are directly connected to the root, flattening
// the tree structure for future operations.
func (d *DSU) Find(x int) int {
	if x < 0 || x >= d.size {
		return -1 // Invalid element
	}

	// Path compression: make every node on the path point directly to the root
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

// Union merges the sets containing elements x and y.
// Implements union by rank optimization: the tree with smaller rank
// is attached under the root of the tree with larger rank.
// This keeps the trees balanced and maintains efficiency.
// Returns true if union was performed (elements were in different sets),
// false if elements were already in the same set.
func (d *DSU) Union(x, y int) bool {
	if x < 0 || x >= d.size || y < 0 || y >= d.size {
		return false // Invalid elements
	}

	rootX := d.Find(x)
	rootY := d.Find(y)

	// Already in the same set
	if rootX == rootY {
		return false
	}

	// Union by rank: attach the tree with smaller rank under the tree with larger rank
	if d.rank[rootX] < d.rank[rootY] {
		d.parent[rootX] = rootY
	} else if d.rank[rootX] > d.rank[rootY] {
		d.parent[rootY] = rootX
	} else {
		// Same rank: make one the parent and increase its rank
		d.parent[rootY] = rootX
		d.rank[rootX]++
	}

	// Decrease the number of components since we merged two sets
	d.components--
	return true
}

// Connected returns true if elements x and y are in the same set.
// This is a convenience method equivalent to Find(x) == Find(y).
func (d *DSU) Connected(x, y int) bool {
	if x < 0 || x >= d.size || y < 0 || y >= d.size {
		return false // Invalid elements are not connected to anything
	}
	return d.Find(x) == d.Find(y)
}

// ComponentCount returns the current number of disjoint sets (connected components).
// This value starts at n (when each element is its own set) and decreases
// by 1 with each successful union operation.
func (d *DSU) ComponentCount() int {
	return d.components
}

// Size returns the total number of elements in the DSU.
// This is the n value that was used during initialization.
func (d *DSU) Size() int {
	return d.size
}
