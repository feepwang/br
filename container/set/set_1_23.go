//go:build go1.23
// +build go1.23

package set

import "iter"

// Set is a map-based implementation of the Interface.
// It uses a map with empty struct values for memory efficiency.
type Set[T comparable] struct {
	data map[T]struct{}
}

// New creates and returns a new empty set.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// NewWithElements creates and returns a new set containing the given elements.
func NewWithElements[T comparable](elements ...T) *Set[T] {
	s := New[T]()
	for _, element := range elements {
		s.Add(element)
	}
	return s
}

// FromSlice creates and returns a new set containing all unique elements from the slice.
func FromSlice[T comparable](slice []T) *Set[T] {
	s := New[T]()
	for _, element := range slice {
		s.Add(element)
	}
	return s
}

// Add inserts an element into the set.
// Returns true if the element was newly added, false if it already existed.
func (s *Set[T]) Add(element T) bool {
	if s == nil {
		return false
	}
	if s.data == nil {
		s.data = make(map[T]struct{})
	}
	
	_, exists := s.data[element]
	if !exists {
		s.data[element] = struct{}{}
	}
	return !exists
}

// Remove deletes an element from the set.
// Returns true if the element was found and removed, false if it didn't exist.
func (s *Set[T]) Remove(element T) bool {
	if s == nil || s.data == nil {
		return false
	}
	
	_, exists := s.data[element]
	if exists {
		delete(s.data, element)
	}
	return exists
}

// Contains checks if an element exists in the set.
func (s *Set[T]) Contains(element T) bool {
	if s == nil || s.data == nil {
		return false
	}
	_, exists := s.data[element]
	return exists
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	if s == nil || s.data == nil {
		return 0
	}
	return len(s.data)
}

// IsEmpty returns true if the set contains no elements.
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	if s.data != nil {
		for k := range s.data {
			delete(s.data, k)
		}
	}
}

// Slice returns all elements as a slice in no particular order.
func (s *Set[T]) Slice() []T {
	if s.data == nil {
		return nil
	}
	
	result := make([]T, 0, len(s.data))
	for element := range s.data {
		result = append(result, element)
	}
	return result
}

// Equal returns true if this set contains exactly the same elements as other.
func (s *Set[T]) Equal(other Interface[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	
	for element := range s.data {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsSubset returns true if all elements in this set are contained in other.
func (s *Set[T]) IsSubset(other Interface[T]) bool {
	for element := range s.data {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if this set contains all elements from other.
func (s *Set[T]) IsSuperset(other Interface[T]) bool {
	return other.IsSubset(s)
}

// Union returns a new set containing all elements from both sets.
func (s *Set[T]) Union(other Interface[T]) Interface[T] {
	result := New[T]()
	
	// Add all elements from this set
	for element := range s.data {
		result.Add(element)
	}
	
	// Add all elements from the other set
	for _, element := range other.Slice() {
		result.Add(element)
	}
	
	return result
}

// Intersection returns a new set containing elements present in both sets.
func (s *Set[T]) Intersection(other Interface[T]) Interface[T] {
	result := New[T]()
	
	// Choose the smaller set to iterate over for better performance
	if s.Len() <= other.Len() {
		for element := range s.data {
			if other.Contains(element) {
				result.Add(element)
			}
		}
	} else {
		for _, element := range other.Slice() {
			if s.Contains(element) {
				result.Add(element)
			}
		}
	}
	
	return result
}

// Difference returns a new set containing elements in this set but not in other.
func (s *Set[T]) Difference(other Interface[T]) Interface[T] {
	result := New[T]()
	
	for element := range s.data {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	
	return result
}

// SymmetricDifference returns a new set containing elements in either set but not in both.
func (s *Set[T]) SymmetricDifference(other Interface[T]) Interface[T] {
	result := New[T]()
	
	// Add elements from this set that are not in other
	for element := range s.data {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	
	// Add elements from other set that are not in this set
	for _, element := range other.Slice() {
		if !s.Contains(element) {
			result.Add(element)
		}
	}
	
	return result
}

// All returns an iterator over all elements in the set.
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for element := range s.data {
			if !yield(element) {
				return
			}
		}
	}
}