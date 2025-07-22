//go:build !go1.23
// +build !go1.23

package union_find

import (
	"reflect"
	"testing"
)

func TestUnionFindBasic(t *testing.T) {
	uf := NewUnionFind(5)

	// Test initial state
	if uf.Size() != 5 {
		t.Errorf("Expected size 5, got %d", uf.Size())
	}
	if uf.Count() != 5 {
		t.Errorf("Expected count 5, got %d", uf.Count())
	}

	// Test initial separation
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if uf.Connected(i, j) {
				t.Errorf("Elements %d and %d should not be connected initially", i, j)
			}
		}
	}
}

func TestUnionFindUnion(t *testing.T) {
	uf := NewUnionFind(5)

	// Union 0 and 1
	uf.Union(0, 1)
	if !uf.Connected(0, 1) {
		t.Error("0 and 1 should be connected after union")
	}
	if uf.Count() != 4 {
		t.Errorf("Expected count 4 after one union, got %d", uf.Count())
	}

	// Union 2 and 3
	uf.Union(2, 3)
	if !uf.Connected(2, 3) {
		t.Error("2 and 3 should be connected after union")
	}
	if uf.Count() != 3 {
		t.Errorf("Expected count 3 after two unions, got %d", uf.Count())
	}

	// Union sets {0,1} and {2,3}
	uf.Union(1, 2)
	if !uf.Connected(0, 3) {
		t.Error("0 and 3 should be connected after union of their sets")
	}
	if uf.Count() != 2 {
		t.Errorf("Expected count 2 after merging sets, got %d", uf.Count())
	}
}

func TestUnionFindFind(t *testing.T) {
	uf := NewUnionFind(5)

	// Initially, each element should be its own root
	for i := 0; i < 5; i++ {
		if uf.Find(i) != i {
			t.Errorf("Element %d should be its own root initially", i)
		}
	}

	// After union, find should return same root for connected elements
	uf.Union(0, 1)
	uf.Union(1, 2)

	root0 := uf.Find(0)
	root1 := uf.Find(1)
	root2 := uf.Find(2)

	if root0 != root1 || root1 != root2 {
		t.Error("Elements 0, 1, 2 should have the same root after unions")
	}
}

func TestUnionFindSetSize(t *testing.T) {
	uf := NewUnionFind(6)

	// Initially, each set should have size 1
	for i := 0; i < 6; i++ {
		if uf.SetSize(i) != 1 {
			t.Errorf("Initial set size for element %d should be 1, got %d", i, uf.SetSize(i))
		}
	}

	// Union elements to form sets of different sizes
	uf.Union(0, 1) // Set {0, 1}: size 2
	uf.Union(2, 3) // Set {2, 3}: size 2
	uf.Union(2, 4) // Set {2, 3, 4}: size 3
	// Element 5 remains alone: size 1

	if uf.SetSize(0) != 2 {
		t.Errorf("Set size for element 0 should be 2, got %d", uf.SetSize(0))
	}
	if uf.SetSize(1) != 2 {
		t.Errorf("Set size for element 1 should be 2, got %d", uf.SetSize(1))
	}
	if uf.SetSize(2) != 3 {
		t.Errorf("Set size for element 2 should be 3, got %d", uf.SetSize(2))
	}
	if uf.SetSize(3) != 3 {
		t.Errorf("Set size for element 3 should be 3, got %d", uf.SetSize(3))
	}
	if uf.SetSize(4) != 3 {
		t.Errorf("Set size for element 4 should be 3, got %d", uf.SetSize(4))
	}
	if uf.SetSize(5) != 1 {
		t.Errorf("Set size for element 5 should be 1, got %d", uf.SetSize(5))
	}
}

func TestUnionFindSets(t *testing.T) {
	uf := NewUnionFind(6)

	// Create some unions
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(2, 4)
	// Element 5 remains alone

	sets := uf.Sets()
	expectedSets := [][]int{
		{0, 1},
		{2, 3, 4},
		{5},
	}

	if len(sets) != len(expectedSets) {
		t.Errorf("Expected %d sets, got %d", len(expectedSets), len(sets))
	}

	// Check if all expected sets are present
	for _, expectedSet := range expectedSets {
		found := false
		for _, actualSet := range sets {
			if reflect.DeepEqual(expectedSet, actualSet) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected set %v not found in result %v", expectedSet, sets)
		}
	}
}

func TestUnionFindReset(t *testing.T) {
	uf := NewUnionFind(5)

	// Create some unions
	uf.Union(0, 1)
	uf.Union(2, 3)

	// Reset to different size
	uf.Reset(3)

	if uf.Size() != 3 {
		t.Errorf("Expected size 3 after reset, got %d", uf.Size())
	}
	if uf.Count() != 3 {
		t.Errorf("Expected count 3 after reset, got %d", uf.Count())
	}

	// All elements should be in separate sets
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			if uf.Connected(i, j) {
				t.Errorf("Elements %d and %d should not be connected after reset", i, j)
			}
		}
	}
}

func TestUnionFindInvalidInputs(t *testing.T) {
	uf := NewUnionFind(5)

	// Test invalid Find
	if uf.Find(-1) != -1 {
		t.Error("Find with negative index should return -1")
	}
	if uf.Find(5) != -1 {
		t.Error("Find with out-of-bounds index should return -1")
	}

	// Test invalid Union (should not panic)
	uf.Union(-1, 0)
	uf.Union(0, 5)
	uf.Union(-1, 5)

	// Test invalid Connected
	if uf.Connected(-1, 0) {
		t.Error("Connected with negative index should return false")
	}
	if uf.Connected(0, 5) {
		t.Error("Connected with out-of-bounds index should return false")
	}

	// Test invalid SetSize
	if uf.SetSize(-1) != 0 {
		t.Error("SetSize with negative index should return 0")
	}
	if uf.SetSize(5) != 0 {
		t.Error("SetSize with out-of-bounds index should return 0")
	}
}

func TestUnionFindZeroSize(t *testing.T) {
	uf := NewUnionFind(0)

	if uf.Size() != 0 {
		t.Errorf("Expected size 0, got %d", uf.Size())
	}
	if uf.Count() != 0 {
		t.Errorf("Expected count 0, got %d", uf.Count())
	}

	sets := uf.Sets()
	if len(sets) != 0 {
		t.Errorf("Expected 0 sets for empty UnionFind, got %d", len(sets))
	}
}

func TestUnionFindNegativeSize(t *testing.T) {
	uf := NewUnionFind(-5)

	if uf.Size() != 0 {
		t.Errorf("Expected size 0 for negative input, got %d", uf.Size())
	}
	if uf.Count() != 0 {
		t.Errorf("Expected count 0 for negative input, got %d", uf.Count())
	}
}

func TestUnionFindPathCompression(t *testing.T) {
	uf := NewUnionFind(10)

	// Create a chain: 0 -> 1 -> 2 -> 3 -> 4
	for i := 0; i < 4; i++ {
		uf.Union(i, i+1)
	}

	// Find should trigger path compression
	root := uf.Find(0)

	// All elements should now point directly to the root
	for i := 0; i < 5; i++ {
		if uf.Find(i) != root {
			t.Errorf("Element %d should have root %d after path compression", i, root)
		}
	}
}

func TestUnionFindComplexScenario(t *testing.T) {
	uf := NewUnionFind(10)

	// Create a complex scenario
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(3, 4)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(8, 9)

	// Merge some sets
	uf.Union(2, 3) // Merge {0,1,2} with {3,4}
	uf.Union(7, 8) // Merge {5,6,7} with {8,9}

	expectedCount := 2 // Two sets: {0,1,2,3,4} and {5,6,7,8,9}
	if uf.Count() != expectedCount {
		t.Errorf("Expected %d sets, got %d", expectedCount, uf.Count())
	}

	// Check connectivity within sets
	group1 := []int{0, 1, 2, 3, 4}
	group2 := []int{5, 6, 7, 8, 9}

	for i := 0; i < len(group1); i++ {
		for j := i + 1; j < len(group1); j++ {
			if !uf.Connected(group1[i], group1[j]) {
				t.Errorf("Elements %d and %d should be connected", group1[i], group1[j])
			}
		}
	}

	for i := 0; i < len(group2); i++ {
		for j := i + 1; j < len(group2); j++ {
			if !uf.Connected(group2[i], group2[j]) {
				t.Errorf("Elements %d and %d should be connected", group2[i], group2[j])
			}
		}
	}

	// Check no connectivity between groups
	for _, x := range group1 {
		for _, y := range group2 {
			if uf.Connected(x, y) {
				t.Errorf("Elements %d and %d should not be connected", x, y)
			}
		}
	}
}

func TestUnionFindInterfaceCompliance(t *testing.T) {
	var uf Interface = NewUnionFind(5)

	// Test that all interface methods are available
	uf.Union(0, 1)
	uf.Find(0)
	uf.Connected(0, 1)
	uf.Count()
	uf.Size()
	uf.SetSize(0)
	uf.Sets()
	uf.Reset(3)
}
