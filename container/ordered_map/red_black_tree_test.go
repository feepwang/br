package ordered_map

import (
	"testing"

	"github.com/feepwang/br/container/pair"
)

func TestRedBlackTreeBasic(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	// Test empty tree
	if tree.Len() != 0 {
		t.Errorf("Expected length 0, got %d", tree.Len())
	}
	if tree.Cap() != 0 {
		t.Errorf("Expected capacity 0, got %d", tree.Cap())
	}

	// Test Get on empty tree
	if _, ok := tree.Get(1); ok {
		t.Error("Expected false when getting from empty tree")
	}

	// Test Has on empty tree
	if tree.Has(1) {
		t.Error("Expected false when checking existence in empty tree")
	}
}

func TestRedBlackTreeInsertAndGet(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	// Insert some values
	tree.Set(5, "five")
	tree.Set(3, "three")
	tree.Set(7, "seven")
	tree.Set(1, "one")
	tree.Set(9, "nine")

	// Check length and capacity
	if tree.Len() != 5 {
		t.Errorf("Expected length 5, got %d", tree.Len())
	}
	if tree.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", tree.Cap())
	}

	// Test Get
	if val, ok := tree.Get(5); !ok || val != "five" {
		t.Errorf("Expected ('five', true), got ('%s', %t)", val, ok)
	}
	if val, ok := tree.Get(1); !ok || val != "one" {
		t.Errorf("Expected ('one', true), got ('%s', %t)", val, ok)
	}
	if _, ok := tree.Get(10); ok {
		t.Error("Expected false when getting non-existent key")
	}

	// Test Has
	if !tree.Has(7) {
		t.Error("Expected true for existing key")
	}
	if tree.Has(10) {
		t.Error("Expected false for non-existent key")
	}
}

func TestRedBlackTreeGetMutable(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Set(1, "original")

	// Test GetMutable
	if ptr, ok := tree.GetMutable(1); !ok || *ptr != "original" {
		t.Errorf("Expected ('original', true), got ('%s', %t)", *ptr, ok)
	}

	// Modify through pointer
	if ptr, ok := tree.GetMutable(1); ok {
		*ptr = "modified"
	}

	// Verify modification
	if val, _ := tree.Get(1); val != "modified" {
		t.Errorf("Expected 'modified', got '%s'", val)
	}

	// Test GetMutable on non-existent key
	if _, ok := tree.GetMutable(99); ok {
		t.Error("Expected false for non-existent key")
	}
}

func TestRedBlackTreeUpdate(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Set(1, "first")

	// Update existing key
	tree.Set(1, "updated")
	if val, _ := tree.Get(1); val != "updated" {
		t.Errorf("Expected 'updated', got '%s'", val)
	}

	// Length should remain the same
	if tree.Len() != 1 {
		t.Errorf("Expected length 1, got %d", tree.Len())
	}
}

func TestRedBlackTreeDelete(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	// Insert some values
	tree.Set(5, "five")
	tree.Set(3, "three")
	tree.Set(7, "seven")
	tree.Set(1, "one")
	tree.Set(9, "nine")

	// Delete existing key
	if !tree.Delete(3) {
		t.Error("Expected true when deleting existing key")
	}
	if tree.Len() != 4 {
		t.Errorf("Expected length 4, got %d", tree.Len())
	}
	if tree.Has(3) {
		t.Error("Expected false after deleting key")
	}

	// Delete non-existent key
	if tree.Delete(99) {
		t.Error("Expected false when deleting non-existent key")
	}
	if tree.Len() != 4 {
		t.Errorf("Length should remain 4, got %d", tree.Len())
	}

	// Delete root
	if !tree.Delete(5) {
		t.Error("Expected true when deleting root")
	}
	if tree.Len() != 3 {
		t.Errorf("Expected length 3, got %d", tree.Len())
	}
}

func TestRedBlackTreeKeysValuesOrder(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	// Insert in random order
	tree.Set(5, "five")
	tree.Set(2, "two")
	tree.Set(8, "eight")
	tree.Set(1, "one")
	tree.Set(7, "seven")

	// Check keys are in order
	keys := tree.Keys()
	expected := []int{1, 2, 5, 7, 8}
	if len(keys) != len(expected) {
		t.Errorf("Expected %d keys, got %d", len(expected), len(keys))
	}
	for i, key := range keys {
		if key != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], key)
		}
	}

	// Check values are in corresponding order
	values := tree.Values()
	expectedValues := []string{"one", "two", "five", "seven", "eight"}
	if len(values) != len(expectedValues) {
		t.Errorf("Expected %d values, got %d", len(expectedValues), len(values))
	}
	for i, value := range values {
		if value != expectedValues[i] {
			t.Errorf("At index %d, expected %s, got %s", i, expectedValues[i], value)
		}
	}
}

func TestRedBlackTreePairs(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Set(3, "three")
	tree.Set(1, "one")
	tree.Set(2, "two")

	pairs := tree.Pairs()
	expected := []pair.Pair[int, string]{
		{First: 1, Second: "one"},
		{First: 2, Second: "two"},
		{First: 3, Second: "three"},
	}

	if len(pairs) != len(expected) {
		t.Errorf("Expected %d pairs, got %d", len(expected), len(pairs))
	}
	for i, p := range pairs {
		if p.First != expected[i].First || p.Second != expected[i].Second {
			t.Errorf("At index %d, expected (%d, %s), got (%d, %s)",
				i, expected[i].First, expected[i].Second, p.First, p.Second)
		}
	}
}

func TestRedBlackTreeInterfaceCompliance(t *testing.T) {
	// This test ensures RedBlackTree implements Interface
	var _ Interface[int, string] = NewRedBlackTree[int, string]()
}
