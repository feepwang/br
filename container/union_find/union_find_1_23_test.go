//go:build go1.23
// +build go1.23

package union_find

import (
	"slices"
	"testing"
)

func TestUnionFind123AllSets(t *testing.T) {
	uf := NewUnionFind(6)

	// Create some unions
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(2, 4)
	// Element 5 remains alone

	var sets [][]int
	for set := range uf.AllSets() {
		sets = append(sets, set)
	}

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
			if slices.Equal(expectedSet, actualSet) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected set %v not found in result %v", expectedSet, sets)
		}
	}
}

func TestUnionFind123SetMembers(t *testing.T) {
	uf := NewUnionFind(6)

	// Create some unions
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(3, 4)
	// Element 5 remains alone

	// Test set containing element 0 (should be {0, 1, 2})
	var members []int
	for member := range uf.SetMembers(0) {
		members = append(members, member)
	}
	slices.Sort(members)
	expected := []int{0, 1, 2}
	if !slices.Equal(members, expected) {
		t.Errorf("Expected set members %v for element 0, got %v", expected, members)
	}

	// Test set containing element 3 (should be {3, 4})
	members = nil
	for member := range uf.SetMembers(3) {
		members = append(members, member)
	}
	slices.Sort(members)
	expected = []int{3, 4}
	if !slices.Equal(members, expected) {
		t.Errorf("Expected set members %v for element 3, got %v", expected, members)
	}

	// Test set containing element 5 (should be {5})
	members = nil
	for member := range uf.SetMembers(5) {
		members = append(members, member)
	}
	expected = []int{5}
	if !slices.Equal(members, expected) {
		t.Errorf("Expected set members %v for element 5, got %v", expected, members)
	}
}

func TestUnionFind123IteratorEarlyTermination(t *testing.T) {
	uf := NewUnionFind(10)

	// Create multiple sets
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)
	uf.Union(6, 7)
	uf.Union(8, 9)

	// Test early termination of AllSets iterator
	count := 0
	for set := range uf.AllSets() {
		count++
		if len(set) == 2 && count == 2 {
			break // Stop after processing 2 sets
		}
	}

	if count != 2 {
		t.Errorf("Expected to process 2 sets before breaking, processed %d", count)
	}

	// Test early termination of SetMembers iterator
	uf.Union(0, 2) // Now {0, 1, 2, 3} are in one set
	memberCount := 0
	for range uf.SetMembers(0) {
		memberCount++
		if memberCount == 2 {
			break // Stop after processing 2 members
		}
	}

	if memberCount != 2 {
		t.Errorf("Expected to process 2 members before breaking, processed %d", memberCount)
	}
}

func TestUnionFind123InvalidSetMembers(t *testing.T) {
	uf := NewUnionFind(5)

	// Test invalid element
	count := 0
	for range uf.SetMembers(-1) {
		count++
	}
	if count != 0 {
		t.Errorf("Expected 0 members for invalid element -1, got %d", count)
	}

	count = 0
	for range uf.SetMembers(5) {
		count++
	}
	if count != 0 {
		t.Errorf("Expected 0 members for out-of-bounds element 5, got %d", count)
	}
}

func TestUnionFind123EmptyUnionFind(t *testing.T) {
	uf := NewUnionFind(0)

	// Test AllSets with empty union-find
	count := 0
	for range uf.AllSets() {
		count++
	}
	if count != 0 {
		t.Errorf("Expected 0 sets for empty union-find, got %d", count)
	}
}

func TestUnionFind123CollectAllSets(t *testing.T) {
	uf := NewUnionFind(8)

	// Create specific pattern
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)
	uf.Union(6, 7)

	// Collect all sets using iterator
	var allSets [][]int
	for set := range uf.AllSets() {
		allSets = append(allSets, slices.Clone(set))
	}

	// Should have 4 sets, each with 2 elements
	if len(allSets) != 4 {
		t.Errorf("Expected 4 sets, got %d", len(allSets))
	}

	for i, set := range allSets {
		if len(set) != 2 {
			t.Errorf("Set %d should have 2 elements, got %d", i, len(set))
		}
	}

	// Verify total elements
	totalElements := 0
	for _, set := range allSets {
		totalElements += len(set)
	}
	if totalElements != 8 {
		t.Errorf("Expected total 8 elements across all sets, got %d", totalElements)
	}
}

func TestUnionFind123ConsistentOrder(t *testing.T) {
	uf := NewUnionFind(6)

	// Create unions
	uf.Union(5, 4)
	uf.Union(4, 3)
	uf.Union(1, 0)

	// Collect sets multiple times to ensure consistent ordering
	var sets1, sets2 [][]int

	for set := range uf.AllSets() {
		sets1 = append(sets1, slices.Clone(set))
	}

	for set := range uf.AllSets() {
		sets2 = append(sets2, slices.Clone(set))
	}

	if len(sets1) != len(sets2) {
		t.Errorf("Set count should be consistent: %d vs %d", len(sets1), len(sets2))
	}

	for i := range sets1 {
		if !slices.Equal(sets1[i], sets2[i]) {
			t.Errorf("Set %d should be consistent: %v vs %v", i, sets1[i], sets2[i])
		}
	}
}

func TestUnionFind123LargeSetIteration(t *testing.T) {
	uf := NewUnionFind(1000)

	// Union all elements into one large set
	for i := 1; i < 1000; i++ {
		uf.Union(0, i)
	}

	// Count members using iterator
	memberCount := 0
	for range uf.SetMembers(500) { // Any element should give the same result
		memberCount++
	}

	if memberCount != 1000 {
		t.Errorf("Expected 1000 members in the set, got %d", memberCount)
	}

	// Verify only one set exists
	setCount := 0
	for range uf.AllSets() {
		setCount++
	}

	if setCount != 1 {
		t.Errorf("Expected 1 set, got %d", setCount)
	}
}
