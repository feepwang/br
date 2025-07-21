//go:build go1.23
// +build go1.23

package set

import (
	"slices"
	"testing"
)

func TestAll_Iterator(t *testing.T) {
	s := NewWithElements(1, 2, 3, 4, 5)
	
	var collected []int
	for element := range s.All() {
		collected = append(collected, element)
	}
	
	// Sort both slices for comparison since iteration order is not guaranteed
	slices.Sort(collected)
	expected := []int{1, 2, 3, 4, 5}
	
	if len(collected) != len(expected) {
		t.Errorf("Expected %d elements, got %d", len(expected), len(collected))
	}
	
	for i, v := range expected {
		if collected[i] != v {
			t.Errorf("At index %d: expected %v, got %v", i, v, collected[i])
		}
	}
}

func TestAll_EmptySet(t *testing.T) {
	s := New[string]()
	
	count := 0
	for range s.All() {
		count++
	}
	
	if count != 0 {
		t.Errorf("Expected 0 iterations for empty set, got %d", count)
	}
}

func TestAll_EarlyReturn(t *testing.T) {
	s := NewWithElements(1, 2, 3, 4, 5)
	
	count := 0
	for element := range s.All() {
		count++
		if element == 3 || count >= 3 { // Stop after finding 3 or after 3 iterations
			break
		}
	}
	
	if count > 3 {
		t.Errorf("Expected at most 3 iterations, got %d", count)
	}
}

func TestAll_StringSet(t *testing.T) {
	s := NewWithElements("apple", "banana", "cherry")
	
	var collected []string
	for element := range s.All() {
		collected = append(collected, element)
	}
	
	if len(collected) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(collected))
	}
	
	// Check that all expected elements are present
	expected := map[string]bool{
		"apple":  false,
		"banana": false,
		"cherry": false,
	}
	
	for _, element := range collected {
		if _, exists := expected[element]; !exists {
			t.Errorf("Unexpected element: %s", element)
		}
		expected[element] = true
	}
	
	for element, found := range expected {
		if !found {
			t.Errorf("Expected element not found: %s", element)
		}
	}
}

func BenchmarkAll_Iterator(b *testing.B) {
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Add(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range s.All() {
			// Iterate through all elements
		}
	}
}