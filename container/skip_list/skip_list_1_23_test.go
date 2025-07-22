//go:build go1.23
// +build go1.23

package skip_list

import (
	"cmp"
	"reflect"
	"testing"

	"github.com/feepwang/br/container/pair"
)

func TestSkipList123Basic(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Test empty skip list
	if sl.Len() != 0 {
		t.Errorf("Expected length 0, got %d", sl.Len())
	}

	// Test iterator on empty skip list
	for k, v := range sl.All() {
		t.Errorf("Expected no elements in empty skip list, got (%d, %s)", k, v)
	}
}

func TestSkipList123SetAndIterateAll(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Insert test data
	testData := map[int]string{
		3: "three",
		1: "one",
		4: "four",
		2: "two",
		5: "five",
	}

	for key, value := range testData {
		sl.Set(key, value)
	}

	// Test All iterator
	var keys []int
	var values []string

	for k, v := range sl.All() {
		keys = append(keys, k)
		values = append(values, v)
	}

	expectedKeys := []int{1, 2, 3, 4, 5}
	expectedValues := []string{"one", "two", "three", "four", "five"}

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}
}

func TestSkipList123AllFrom(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Insert test data
	for i := 1; i <= 10; i++ {
		sl.Set(i, string(rune('A'+i-1)))
	}

	// Test AllFrom iterator
	var keys []int
	var values []string

	for k, v := range sl.AllFrom(5) {
		keys = append(keys, k)
		values = append(values, v)
	}

	expectedKeys := []int{5, 6, 7, 8, 9, 10}
	expectedValues := []string{"E", "F", "G", "H", "I", "J"}

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}

	// Test AllFrom with non-existing start key
	keys = nil
	values = nil

	for k, v := range sl.AllFrom(7) {
		keys = append(keys, k)
		values = append(values, v)
		if len(keys) >= 2 { // Limit to 2 elements
			break
		}
	}

	expectedKeysLimited := []int{7, 8}
	expectedValuesLimited := []string{"G", "H"}

	if !reflect.DeepEqual(keys, expectedKeysLimited) {
		t.Errorf("Expected keys %v with break, got %v", expectedKeysLimited, keys)
	}

	if !reflect.DeepEqual(values, expectedValuesLimited) {
		t.Errorf("Expected values %v with break, got %v", expectedValuesLimited, values)
	}
}

func TestSkipList123AllBetween(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Insert test data
	for i := 1; i <= 10; i++ {
		sl.Set(i, string(rune('A'+i-1)))
	}

	// Test AllBetween iterator
	var keys []int
	var values []string

	for k, v := range sl.AllBetween(3, 7) {
		keys = append(keys, k)
		values = append(values, v)
	}

	expectedKeys := []int{3, 4, 5, 6, 7}
	expectedValues := []string{"C", "D", "E", "F", "G"}

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}

	// Test AllBetween with single key range
	keys = nil
	values = nil

	for k, v := range sl.AllBetween(5, 5) {
		keys = append(keys, k)
		values = append(values, v)
	}

	expectedKeysSingle := []int{5}
	expectedValuesSingle := []string{"E"}

	if !reflect.DeepEqual(keys, expectedKeysSingle) {
		t.Errorf("Expected keys %v for single key range, got %v", expectedKeysSingle, keys)
	}

	if !reflect.DeepEqual(values, expectedValuesSingle) {
		t.Errorf("Expected values %v for single key range, got %v", expectedValuesSingle, values)
	}

	// Test AllBetween with range outside existing keys
	keys = nil
	values = nil

	for k, v := range sl.AllBetween(15, 20) {
		keys = append(keys, k)
		values = append(values, v)
	}

	if len(keys) != 0 {
		t.Errorf("Expected no keys for range outside existing keys, got %v", keys)
	}
}

func TestSkipList123IteratorEarlyTermination(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Insert test data
	for i := 1; i <= 10; i++ {
		sl.Set(i, string(rune('A'+i-1)))
	}

	// Test early termination in All iterator
	var keys []int
	for k, _ := range sl.All() {
		keys = append(keys, k)
		if k >= 3 { // Stop after key 3
			break
		}
	}

	expectedKeys := []int{1, 2, 3}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v with early termination, got %v", expectedKeys, keys)
	}

	// Test early termination in AllBetween iterator
	keys = nil
	for k, _ := range sl.AllBetween(2, 8) {
		keys = append(keys, k)
		if len(keys) >= 3 { // Stop after 3 items
			break
		}
	}

	expectedKeysLimited := []int{2, 3, 4}
	if !reflect.DeepEqual(keys, expectedKeysLimited) {
		t.Errorf("Expected keys %v with limited collection, got %v", expectedKeysLimited, keys)
	}
}

func TestSkipList123CollectSlices(t *testing.T) {
	sl := NewOrderedSkipList[string, int]()

	// Insert test data
	testData := map[string]int{
		"banana": 2,
		"apple":  1,
		"cherry": 3,
		"date":   4,
	}

	for key, value := range testData {
		sl.Set(key, value)
	}

	// Use slices.Collect to convert iterator to slice of pairs
	var pairs []pair.Pair[string, int]
	for k, v := range sl.All() {
		pairs = append(pairs, pair.Pair[string, int]{First: k, Second: v})
	}

	expectedPairs := []pair.Pair[string, int]{
		{First: "apple", Second: 1},
		{First: "banana", Second: 2},
		{First: "cherry", Second: 3},
		{First: "date", Second: 4},
	}

	if !reflect.DeepEqual(pairs, expectedPairs) {
		t.Errorf("Expected pairs %v, got %v", expectedPairs, pairs)
	}

	// Test collecting just keys using a helper function
	var collectedKeys []string
	for k, _ := range sl.AllFrom("banana") {
		collectedKeys = append(collectedKeys, k)
	}

	expectedKeys := []string{"banana", "cherry", "date"}
	if !reflect.DeepEqual(collectedKeys, expectedKeys) {
		t.Errorf("Expected keys %v from banana, got %v", expectedKeys, collectedKeys)
	}
}

func TestSkipList123CustomComparator(t *testing.T) {
	// Create skip list with custom reverse order comparator
	sl := NewSkipList[int, string](func(a, b int) int {
		return cmp.Compare(b, a) // Reverse order
	})

	// Insert test data
	for i := 1; i <= 5; i++ {
		sl.Set(i, string(rune('A'+i-1)))
	}

	// Keys should be in reverse order
	var keys []int
	for k, _ := range sl.All() {
		keys = append(keys, k)
	}

	expectedKeys := []int{5, 4, 3, 2, 1} // Reverse order
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v in reverse order, got %v", expectedKeys, keys)
	}

	// Test range operations with custom comparator
	var rangeKeys []int
	for k, _ := range sl.AllBetween(2, 4) {
		rangeKeys = append(rangeKeys, k)
	}

	expectedRangeKeys := []int{4, 3, 2} // In reverse order within range
	if !reflect.DeepEqual(rangeKeys, expectedRangeKeys) {
		t.Errorf("Expected range keys %v, got %v", expectedRangeKeys, rangeKeys)
	}
}

func TestSkipList123StringComparison(t *testing.T) {
	sl := NewOrderedSkipList[string, int]()

	// Insert words
	words := []string{"zebra", "apple", "banana", "orange", "grape"}
	for i, word := range words {
		sl.Set(word, i)
	}

	// Collect all keys using iterator
	var keys []string
	for k, _ := range sl.All() {
		keys = append(keys, k)
	}

	// Should be in lexicographic order
	expectedKeys := []string{"apple", "banana", "grape", "orange", "zebra"}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v in lexicographic order, got %v", expectedKeys, keys)
	}

	// Test prefix-like range
	var prefixKeys []string
	for k, _ := range sl.AllBetween("b", "g") {
		prefixKeys = append(prefixKeys, k)
	}

	expectedPrefixKeys := []string{"banana"} // "grape" > "g" so it's not included
	if !reflect.DeepEqual(prefixKeys, expectedPrefixKeys) {
		t.Errorf("Expected prefix keys %v, got %v", expectedPrefixKeys, prefixKeys)
	}
}

func TestSkipList123ComplexIteratorScenario(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Insert large dataset
	for i := 0; i < 100; i += 2 { // Even numbers only
		sl.Set(i, string(rune('A'+(i%26))))
	}

	// Count elements using All iterator
	count := 0
	for range sl.All() {
		count++
	}

	if count != 50 {
		t.Errorf("Expected count 50, got %d", count)
	}

	// Count elements in specific range using AllBetween
	rangeCount := 0
	for range sl.AllBetween(20, 40) {
		rangeCount++
	}

	expectedRangeCount := 11 // 20, 22, 24, ..., 40
	if rangeCount != expectedRangeCount {
		t.Errorf("Expected range count %d, got %d", expectedRangeCount, rangeCount)
	}

	// Test AllFrom with value beyond existing range
	beyondCount := 0
	for range sl.AllFrom(200) {
		beyondCount++
	}

	if beyondCount != 0 {
		t.Errorf("Expected no elements beyond range, got %d", beyondCount)
	}

	// Modify skip list during iteration preparation (not during iteration itself)
	// Add odd numbers
	for i := 1; i < 100; i += 2 {
		sl.Set(i, string(rune('a'+(i%26))))
	}

	// Now count all elements
	allCount := 0
	for range sl.All() {
		allCount++
	}

	if allCount != 100 {
		t.Errorf("Expected total count 100, got %d", allCount)
	}

	// Test mixed even/odd retrieval using iterator
	var evenOddPattern []bool
	count = 0
	for k, _ := range sl.AllBetween(0, 9) {
		evenOddPattern = append(evenOddPattern, k%2 == 0)
		count++
		if count >= 10 {
			break
		}
	}

	expectedPattern := []bool{true, false, true, false, true, false, true, false, true, false}
	if !reflect.DeepEqual(evenOddPattern, expectedPattern) {
		t.Errorf("Expected even/odd pattern %v, got %v", expectedPattern, evenOddPattern)
	}
}
