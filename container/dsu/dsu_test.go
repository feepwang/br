package dsu

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		wantNil  bool
		wantSize int
		wantComp int
	}{
		{
			name:     "valid size",
			n:        5,
			wantNil:  false,
			wantSize: 5,
			wantComp: 5,
		},
		{
			name:    "zero size",
			n:       0,
			wantNil: true,
		},
		{
			name:    "negative size",
			n:       -1,
			wantNil: true,
		},
		{
			name:     "single element",
			n:        1,
			wantNil:  false,
			wantSize: 1,
			wantComp: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsu := NewDSU(tt.n)
			if tt.wantNil {
				if dsu != nil {
					t.Errorf("New(%d) = %v, want nil", tt.n, dsu)
				}
				return
			}

			if dsu == nil {
				t.Fatalf("New(%d) = nil, want non-nil", tt.n)
			}

			if got := dsu.Size(); got != tt.wantSize {
				t.Errorf("Size() = %d, want %d", got, tt.wantSize)
			}

			if got := dsu.ComponentCount(); got != tt.wantComp {
				t.Errorf("ComponentCount() = %d, want %d", got, tt.wantComp)
			}
		})
	}
}

func TestFind(t *testing.T) {
	dsu := NewDSU(5)
	if dsu == nil {
		t.Fatal("Failed to create DSU")
	}

	// Initially, each element should be its own parent
	for i := 0; i < 5; i++ {
		if got := dsu.Find(i); got != i {
			t.Errorf("Find(%d) = %d, want %d", i, got, i)
		}
	}

	// Test invalid indices
	if got := dsu.Find(-1); got != -1 {
		t.Errorf("Find(-1) = %d, want -1", got)
	}
	if got := dsu.Find(5); got != -1 {
		t.Errorf("Find(5) = %d, want -1", got)
	}
}

func TestUnion(t *testing.T) {
	dsu := NewDSU(5)
	if dsu == nil {
		t.Fatal("Failed to create DSU")
	}

	// Test successful union
	if !dsu.Union(0, 1) {
		t.Error("Union(0, 1) = false, want true")
	}

	// Elements 0 and 1 should now have the same root
	root0 := dsu.Find(0)
	root1 := dsu.Find(1)
	if root0 != root1 {
		t.Errorf("After Union(0, 1): Find(0) = %d, Find(1) = %d, want same root", root0, root1)
	}

	// Component count should decrease
	if got := dsu.ComponentCount(); got != 4 {
		t.Errorf("ComponentCount() = %d, want 4", got)
	}

	// Test union of already connected elements
	if dsu.Union(0, 1) {
		t.Error("Union(0, 1) = true, want false (already connected)")
	}

	// Component count should remain the same
	if got := dsu.ComponentCount(); got != 4 {
		t.Errorf("ComponentCount() = %d, want 4", got)
	}

	// Test invalid indices
	if dsu.Union(-1, 0) {
		t.Error("Union(-1, 0) = true, want false")
	}
	if dsu.Union(0, 5) {
		t.Error("Union(0, 5) = true, want false")
	}
}

func TestConnected(t *testing.T) {
	dsu := NewDSU(5)
	if dsu == nil {
		t.Fatal("Failed to create DSU")
	}

	// Initially, no elements should be connected
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if dsu.Connected(i, j) {
				t.Errorf("Connected(%d, %d) = true, want false", i, j)
			}
		}
	}

	// Each element should be connected to itself
	for i := 0; i < 5; i++ {
		if !dsu.Connected(i, i) {
			t.Errorf("Connected(%d, %d) = false, want true", i, i)
		}
	}

	// Union elements and test connectivity
	dsu.Union(0, 1)
	if !dsu.Connected(0, 1) {
		t.Error("Connected(0, 1) = false, want true after union")
	}
	if !dsu.Connected(1, 0) {
		t.Error("Connected(1, 0) = false, want true after union")
	}

	// Test transitive connectivity
	dsu.Union(1, 2)
	if !dsu.Connected(0, 2) {
		t.Error("Connected(0, 2) = false, want true (transitive)")
	}

	// Test elements not connected
	if dsu.Connected(0, 3) {
		t.Error("Connected(0, 3) = true, want false")
	}

	// Test invalid indices
	if dsu.Connected(-1, 0) {
		t.Error("Connected(-1, 0) = true, want false")
	}
	if dsu.Connected(0, 5) {
		t.Error("Connected(0, 5) = true, want false")
	}
}

func TestComplexOperations(t *testing.T) {
	dsu := NewDSU(10)
	if dsu == nil {
		t.Fatal("Failed to create DSU")
	}

	// Create several connected components:
	// Component 1: {0, 1, 2}
	// Component 2: {3, 4}
	// Component 3: {5, 6, 7}
	// Component 4: {8}
	// Component 5: {9}

	// Build component 1
	dsu.Union(0, 1)
	dsu.Union(1, 2)

	// Build component 2
	dsu.Union(3, 4)

	// Build component 3
	dsu.Union(5, 6)
	dsu.Union(6, 7)

	// Verify component count
	if got := dsu.ComponentCount(); got != 5 {
		t.Errorf("ComponentCount() = %d, want 5", got)
	}

	// Verify connectivity within components
	testCases := []struct {
		x, y      int
		connected bool
	}{
		{0, 1, true}, {0, 2, true}, {1, 2, true}, // Component 1
		{3, 4, true},                             // Component 2
		{5, 6, true}, {5, 7, true}, {6, 7, true}, // Component 3
		{8, 8, true}, // Component 4 (self)
		{9, 9, true}, // Component 5 (self)
		// Cross-component (should not be connected)
		{0, 3, false}, {0, 5, false}, {0, 8, false}, {0, 9, false},
		{3, 5, false}, {3, 8, false}, {3, 9, false},
		{5, 8, false}, {5, 9, false},
		{8, 9, false},
	}

	for _, tc := range testCases {
		if got := dsu.Connected(tc.x, tc.y); got != tc.connected {
			t.Errorf("Connected(%d, %d) = %t, want %t", tc.x, tc.y, got, tc.connected)
		}
	}

	// Merge two components and verify
	dsu.Union(2, 3) // Merge component 1 and 2
	if got := dsu.ComponentCount(); got != 4 {
		t.Errorf("After merging components: ComponentCount() = %d, want 4", got)
	}

	// Now 0-4 should all be connected
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			if !dsu.Connected(i, j) {
				t.Errorf("Connected(%d, %d) = false, want true after component merge", i, j)
			}
		}
	}
}

func TestPathCompression(t *testing.T) {
	dsu := NewDSU(5)
	if dsu == nil {
		t.Fatal("Failed to create DSU")
	}

	// Create a chain: 0 -> 1 -> 2 -> 3 -> 4
	// This creates a worst-case tree structure
	for i := 0; i < 4; i++ {
		dsu.Union(i, i+1)
	}

	// Verify all elements are connected
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !dsu.Connected(i, j) {
				t.Errorf("Connected(%d, %d) = false, want true", i, j)
			}
		}
	}

	// The find operation should perform path compression
	// After calling Find on each element, they should all point more directly to the root
	roots := make([]int, 5)
	for i := 0; i < 5; i++ {
		roots[i] = dsu.Find(i)
	}

	// All roots should be the same
	for i := 1; i < 5; i++ {
		if roots[i] != roots[0] {
			t.Errorf("Find(%d) = %d, want %d (same root as Find(0))", i, roots[i], roots[0])
		}
	}
}

func TestSingleElement(t *testing.T) {
	dsu := NewDSU(1)
	if dsu == nil {
		t.Fatal("Failed to create DSU")
	}

	if got := dsu.Size(); got != 1 {
		t.Errorf("Size() = %d, want 1", got)
	}

	if got := dsu.ComponentCount(); got != 1 {
		t.Errorf("ComponentCount() = %d, want 1", got)
	}

	if got := dsu.Find(0); got != 0 {
		t.Errorf("Find(0) = %d, want 0", got)
	}

	if !dsu.Connected(0, 0) {
		t.Error("Connected(0, 0) = false, want true")
	}

	// Union with itself should return false (already connected)
	if dsu.Union(0, 0) {
		t.Error("Union(0, 0) = true, want false")
	}

	if got := dsu.ComponentCount(); got != 1 {
		t.Errorf("ComponentCount() after Union(0, 0) = %d, want 1", got)
	}
}

// Benchmark tests for performance analysis
func BenchmarkFind(b *testing.B) {
	dsu := NewDSU(1000)
	// Create a long chain to test path compression
	for i := 0; i < 999; i++ {
		dsu.Union(i, i+1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dsu.Find(i % 1000)
	}
}

func BenchmarkUnion(b *testing.B) {
	dsu := NewDSU(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := i % 1000
		y := (i + 1) % 1000
		dsu.Union(x, y)
	}
}

func BenchmarkConnected(b *testing.B) {
	dsu := NewDSU(1000)
	// Create some connections
	for i := 0; i < 500; i++ {
		dsu.Union(i, i+500)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := i % 1000
		y := (i + 1) % 1000
		dsu.Connected(x, y)
	}
}
