package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Fatal("New() returned nil")
	}
	if s.Len() != 0 {
		t.Errorf("New set should be empty, got length %d", s.Len())
	}
	if !s.IsEmpty() {
		t.Error("New set should be empty")
	}
}

func TestNewWithElements(t *testing.T) {
	s := NewWithElements(1, 2, 3, 2, 1) // Include duplicates
	expectedLen := 3
	if s.Len() != expectedLen {
		t.Errorf("Expected length %d, got %d", expectedLen, s.Len())
	}
	
	expected := []int{1, 2, 3}
	for _, v := range expected {
		if !s.Contains(v) {
			t.Errorf("Set should contain %v", v)
		}
	}
}

func TestFromSlice(t *testing.T) {
	slice := []string{"a", "b", "c", "b", "a"}
	s := FromSlice(slice)
	
	expectedLen := 3
	if s.Len() != expectedLen {
		t.Errorf("Expected length %d, got %d", expectedLen, s.Len())
	}
	
	expected := []string{"a", "b", "c"}
	for _, v := range expected {
		if !s.Contains(v) {
			t.Errorf("Set should contain %v", v)
		}
	}
}

func TestAdd(t *testing.T) {
	s := New[int]()
	
	// Test adding new element
	added := s.Add(1)
	if !added {
		t.Error("Add should return true for new element")
	}
	if !s.Contains(1) {
		t.Error("Set should contain added element")
	}
	if s.Len() != 1 {
		t.Errorf("Expected length 1, got %d", s.Len())
	}
	
	// Test adding duplicate element
	added = s.Add(1)
	if added {
		t.Error("Add should return false for duplicate element")
	}
	if s.Len() != 1 {
		t.Errorf("Length should remain 1, got %d", s.Len())
	}
}

func TestRemove(t *testing.T) {
	s := NewWithElements(1, 2, 3)
	
	// Test removing existing element
	removed := s.Remove(2)
	if !removed {
		t.Error("Remove should return true for existing element")
	}
	if s.Contains(2) {
		t.Error("Set should not contain removed element")
	}
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}
	
	// Test removing non-existing element
	removed = s.Remove(4)
	if removed {
		t.Error("Remove should return false for non-existing element")
	}
	if s.Len() != 2 {
		t.Errorf("Length should remain 2, got %d", s.Len())
	}
}

func TestContains(t *testing.T) {
	s := NewWithElements("apple", "banana", "cherry")
	
	if !s.Contains("apple") {
		t.Error("Set should contain 'apple'")
	}
	if s.Contains("orange") {
		t.Error("Set should not contain 'orange'")
	}
}

func TestClear(t *testing.T) {
	s := NewWithElements(1, 2, 3, 4, 5)
	s.Clear()
	
	if !s.IsEmpty() {
		t.Error("Set should be empty after Clear()")
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}
}

func TestSlice(t *testing.T) {
	original := []int{3, 1, 4, 1, 5, 9, 2, 6}
	s := FromSlice(original)
	
	result := s.Slice()
	sort.Ints(result) // Sort for consistent comparison
	
	expected := []int{1, 2, 3, 4, 5, 6, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestEqual(t *testing.T) {
	s1 := NewWithElements(1, 2, 3)
	s2 := NewWithElements(3, 2, 1) // Different order
	s3 := NewWithElements(1, 2, 3, 4)
	
	if !s1.Equal(s2) {
		t.Error("Sets with same elements should be equal")
	}
	if s1.Equal(s3) {
		t.Error("Sets with different elements should not be equal")
	}
}

func TestIsSubset(t *testing.T) {
	s1 := NewWithElements(1, 2)
	s2 := NewWithElements(1, 2, 3, 4)
	s3 := NewWithElements(1, 5)
	
	if !s1.IsSubset(s2) {
		t.Error("s1 should be a subset of s2")
	}
	if s1.IsSubset(s3) {
		t.Error("s1 should not be a subset of s3")
	}
	if !s1.IsSubset(s1) {
		t.Error("Set should be a subset of itself")
	}
}

func TestIsSuperset(t *testing.T) {
	s1 := NewWithElements(1, 2, 3, 4)
	s2 := NewWithElements(1, 2)
	s3 := NewWithElements(1, 5)
	
	if !s1.IsSuperset(s2) {
		t.Error("s1 should be a superset of s2")
	}
	if s1.IsSuperset(s3) {
		t.Error("s1 should not be a superset of s3")
	}
	if !s1.IsSuperset(s1) {
		t.Error("Set should be a superset of itself")
	}
}

func TestUnion(t *testing.T) {
	s1 := NewWithElements(1, 2, 3)
	s2 := NewWithElements(3, 4, 5)
	
	union := s1.Union(s2)
	expected := []int{1, 2, 3, 4, 5}
	
	if union.Len() != len(expected) {
		t.Errorf("Expected union length %d, got %d", len(expected), union.Len())
	}
	
	for _, v := range expected {
		if !union.Contains(v) {
			t.Errorf("Union should contain %v", v)
		}
	}
}

func TestIntersection(t *testing.T) {
	s1 := NewWithElements(1, 2, 3, 4)
	s2 := NewWithElements(3, 4, 5, 6)
	
	intersection := s1.Intersection(s2)
	expected := []int{3, 4}
	
	if intersection.Len() != len(expected) {
		t.Errorf("Expected intersection length %d, got %d", len(expected), intersection.Len())
	}
	
	for _, v := range expected {
		if !intersection.Contains(v) {
			t.Errorf("Intersection should contain %v", v)
		}
	}
	
	// Test empty intersection
	s3 := NewWithElements(7, 8, 9)
	emptyIntersection := s1.Intersection(s3)
	if !emptyIntersection.IsEmpty() {
		t.Error("Intersection of disjoint sets should be empty")
	}
}

func TestDifference(t *testing.T) {
	s1 := NewWithElements(1, 2, 3, 4)
	s2 := NewWithElements(3, 4, 5, 6)
	
	difference := s1.Difference(s2)
	expected := []int{1, 2}
	
	if difference.Len() != len(expected) {
		t.Errorf("Expected difference length %d, got %d", len(expected), difference.Len())
	}
	
	for _, v := range expected {
		if !difference.Contains(v) {
			t.Errorf("Difference should contain %v", v)
		}
	}
	
	// Test difference with no common elements
	s3 := NewWithElements(7, 8, 9)
	fullDifference := s1.Difference(s3)
	if !s1.Equal(fullDifference) {
		t.Error("Difference with disjoint set should equal original set")
	}
}

func TestSymmetricDifference(t *testing.T) {
	s1 := NewWithElements(1, 2, 3, 4)
	s2 := NewWithElements(3, 4, 5, 6)
	
	symDiff := s1.SymmetricDifference(s2)
	expected := []int{1, 2, 5, 6}
	
	if symDiff.Len() != len(expected) {
		t.Errorf("Expected symmetric difference length %d, got %d", len(expected), symDiff.Len())
	}
	
	for _, v := range expected {
		if !symDiff.Contains(v) {
			t.Errorf("Symmetric difference should contain %v", v)
		}
	}
	
	// Test with identical sets (should be empty)
	emptySymDiff := s1.SymmetricDifference(s1)
	if !emptySymDiff.IsEmpty() {
		t.Error("Symmetric difference of identical sets should be empty")
	}
}

func TestSetOperationsWithDifferentTypes(t *testing.T) {
	// Test with strings
	s1 := NewWithElements("apple", "banana", "cherry")
	s2 := NewWithElements("banana", "cherry", "date")
	
	union := s1.Union(s2)
	if union.Len() != 4 {
		t.Errorf("Expected union length 4, got %d", union.Len())
	}
	
	intersection := s1.Intersection(s2)
	if intersection.Len() != 2 {
		t.Errorf("Expected intersection length 2, got %d", intersection.Len())
	}
	
	if !intersection.Contains("banana") || !intersection.Contains("cherry") {
		t.Error("Intersection should contain 'banana' and 'cherry'")
	}
}

func TestNilSetHandling(t *testing.T) {
	var s *Set[int]
	
	// These operations should handle nil receiver gracefully
	if s.Len() != 0 {
		t.Error("Nil set should have length 0")
	}
	if !s.IsEmpty() {
		t.Error("Nil set should be empty")
	}
	if s.Contains(1) {
		t.Error("Nil set should not contain any elements")
	}
	if s.Remove(1) {
		t.Error("Remove from nil set should return false")
	}
	
	// Add should return false for nil set (can't modify nil)
	if s.Add(1) {
		t.Error("Add to nil set should return false")
	}
	if s.Len() != 0 {
		t.Errorf("Nil set should still have length 0, got %d", s.Len())
	}
}

// Benchmark tests
func BenchmarkSetAdd(b *testing.B) {
	s := New[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkSetContains(b *testing.B) {
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(i % 1000)
	}
}

func BenchmarkSetUnion(b *testing.B) {
	s1 := New[int]()
	s2 := New[int]()
	for i := 0; i < 500; i++ {
		s1.Add(i)
		s2.Add(i + 250) // 50% overlap
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.Union(s2)
	}
}

func BenchmarkSetIntersection(b *testing.B) {
	s1 := New[int]()
	s2 := New[int]()
	for i := 0; i < 500; i++ {
		s1.Add(i)
		s2.Add(i + 250) // 50% overlap
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.Intersection(s2)
	}
}