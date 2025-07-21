package set_test

import (
	"fmt"
	"sort"

	"github.com/feepwang/br/container/set"
)

// Example demonstrates basic set operations including union and difference.
func Example() {
	// Create sets with some initial data
	set1 := set.NewWithElements(1, 2, 3, 4)
	set2 := set.NewWithElements(3, 4, 5, 6)

	fmt.Printf("Set 1: %v\n", sorted(set1.Slice()))
	fmt.Printf("Set 2: %v\n", sorted(set2.Slice()))

	// Union: elements in either set
	union := set1.Union(set2)
	fmt.Printf("Union: %v\n", sorted(union.Slice()))

	// Intersection: elements in both sets
	intersection := set1.Intersection(set2)
	fmt.Printf("Intersection: %v\n", sorted(intersection.Slice()))

	// Difference: elements in set1 but not in set2
	difference := set1.Difference(set2)
	fmt.Printf("Difference (1-2): %v\n", sorted(difference.Slice()))

	// Symmetric difference: elements in either set but not in both
	symDiff := set1.SymmetricDifference(set2)
	fmt.Printf("Symmetric Difference: %v\n", sorted(symDiff.Slice()))

	// Output:
	// Set 1: [1 2 3 4]
	// Set 2: [3 4 5 6]
	// Union: [1 2 3 4 5 6]
	// Intersection: [3 4]
	// Difference (1-2): [1 2]
	// Symmetric Difference: [1 2 5 6]
}

// Example_stringSet demonstrates set operations with strings.
func Example_stringSet() {
	// Create sets of programming languages
	frontend := set.NewWithElements("JavaScript", "TypeScript", "HTML", "CSS")
	backend := set.NewWithElements("Go", "JavaScript", "Python", "Java")

	fmt.Printf("Frontend: %v\n", sorted(frontend.Slice()))
	fmt.Printf("Backend: %v\n", sorted(backend.Slice()))

	// Languages used in both frontend and backend
	fullStack := frontend.Intersection(backend)
	fmt.Printf("Full-stack languages: %v\n", sorted(fullStack.Slice()))

	// All languages
	allLanguages := frontend.Union(backend)
	fmt.Printf("All languages: %v\n", sorted(allLanguages.Slice()))

	// Frontend-only languages
	frontendOnly := frontend.Difference(backend)
	fmt.Printf("Frontend-only: %v\n", sorted(frontendOnly.Slice()))

	// Output:
	// Frontend: [CSS HTML JavaScript TypeScript]
	// Backend: [Go Java JavaScript Python]
	// Full-stack languages: [JavaScript]
	// All languages: [CSS Go HTML Java JavaScript Python TypeScript]
	// Frontend-only: [CSS HTML TypeScript]
}

// Example_setOperations demonstrates advanced set operations and relationships.
func Example_setOperations() {
	// Create some test sets
	numbers := set.NewWithElements(1, 2, 3, 4, 5)
	evenNumbers := set.NewWithElements(2, 4, 6, 8)
	smallNumbers := set.NewWithElements(1, 2, 3)

	fmt.Printf("Numbers: %v\n", sorted(numbers.Slice()))
	fmt.Printf("Even numbers: %v\n", sorted(evenNumbers.Slice()))
	fmt.Printf("Small numbers: %v\n", sorted(smallNumbers.Slice()))

	// Check relationships
	fmt.Printf("Small numbers ⊆ Numbers: %t\n", smallNumbers.IsSubset(numbers))
	fmt.Printf("Numbers ⊇ Small numbers: %t\n", numbers.IsSuperset(smallNumbers))
	fmt.Printf("Even numbers ⊆ Numbers: %t\n", evenNumbers.IsSubset(numbers))

	// Set equality
	duplicate := set.NewWithElements(1, 2, 3, 4, 5)
	fmt.Printf("Numbers == Duplicate: %t\n", numbers.Equal(duplicate))

	// Basic operations
	fmt.Printf("Contains 3: %t\n", numbers.Contains(3))
	fmt.Printf("Contains 7: %t\n", numbers.Contains(7))
	fmt.Printf("Size: %d\n", numbers.Len())
	fmt.Printf("Is empty: %t\n", numbers.IsEmpty())

	// Output:
	// Numbers: [1 2 3 4 5]
	// Even numbers: [2 4 6 8]
	// Small numbers: [1 2 3]
	// Small numbers ⊆ Numbers: true
	// Numbers ⊇ Small numbers: true
	// Even numbers ⊆ Numbers: false
	// Numbers == Duplicate: true
	// Contains 3: true
	// Contains 7: false
	// Size: 5
	// Is empty: false
}

// Example_dynamicOperations shows dynamic set manipulation.
func Example_dynamicOperations() {
	// Start with an empty set
	s := set.New[string]()
	fmt.Printf("Initial set: %v (empty: %t)\n", s.Slice(), s.IsEmpty())

	// Add elements
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")
	fmt.Printf("After adding fruits: %v\n", sorted(s.Slice()))

	// Try to add duplicate
	added := s.Add("apple")
	fmt.Printf("Adding duplicate 'apple': %t\n", added)
	fmt.Printf("Set after duplicate: %v\n", sorted(s.Slice()))

	// Remove element
	removed := s.Remove("banana")
	fmt.Printf("Removing 'banana': %t\n", removed)
	fmt.Printf("Set after removal: %v\n", sorted(s.Slice()))

	// Try to remove non-existing element
	removed = s.Remove("grape")
	fmt.Printf("Removing non-existing 'grape': %t\n", removed)

	// Clear the set
	s.Clear()
	fmt.Printf("After clearing: %v (empty: %t)\n", s.Slice(), s.IsEmpty())

	// Output:
	// Initial set: [] (empty: true)
	// After adding fruits: [apple banana cherry]
	// Adding duplicate 'apple': false
	// Set after duplicate: [apple banana cherry]
	// Removing 'banana': true
	// Set after removal: [apple cherry]
	// Removing non-existing 'grape': false
	// After clearing: [] (empty: true)
}

// sorted is a helper function to sort slices for consistent output in examples.
func sorted[T any](slice []T) []T {
	switch v := any(slice).(type) {
	case []int:
		result := make([]int, len(v))
		copy(result, v)
		sort.Ints(result)
		return any(result).([]T)
	case []string:
		result := make([]string, len(v))
		copy(result, v)
		sort.Strings(result)
		return any(result).([]T)
	default:
		return slice
	}
}