package skip_list

import (
	"reflect"
	"testing"

	"github.com/feepwang/br/container/pair"
)

func TestSkipListBasic(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Test empty skip list
	if sl.Len() != 0 {
		t.Errorf("Expected length 0, got %d", sl.Len())
	}

	// Test Get on empty skip list
	if value, exists := sl.Get(1); exists {
		t.Errorf("Expected false when getting from empty skip list, got %v", value)
	}

	// Test Has on empty skip list
	if sl.Has(1) {
		t.Error("Expected false when checking existence in empty skip list")
	}

	// Test Keys on empty skip list
	keys := sl.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected empty slice, got %v", keys)
	}

	// Test Values on empty skip list
	values := sl.Values()
	if len(values) != 0 {
		t.Errorf("Expected empty slice, got %v", values)
	}

	// Test Pairs on empty skip list
	pairs := sl.Pairs()
	if len(pairs) != 0 {
		t.Errorf("Expected empty slice, got %v", pairs)
	}
}

func TestSkipListSetAndGet(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Test setting values
	testData := map[int]string{
		1: "one",
		3: "three",
		2: "two",
		5: "five",
		4: "four",
	}

	for key, value := range testData {
		sl.Set(key, value)
	}

	// Test length
	if sl.Len() != len(testData) {
		t.Errorf("Expected length %d, got %d", len(testData), sl.Len())
	}

	// Test getting existing values
	for key, expectedValue := range testData {
		if value, exists := sl.Get(key); !exists || value != expectedValue {
			t.Errorf("Expected (%s, true) for key %d, got (%s, %t)", expectedValue, key, value, exists)
		}
	}

	// Test Has for existing keys
	for key := range testData {
		if !sl.Has(key) {
			t.Errorf("Expected true for Has(%d)", key)
		}
	}

	// Test getting non-existing values
	nonExistingKeys := []int{0, 6, 10, -1}
	for _, key := range nonExistingKeys {
		if value, exists := sl.Get(key); exists {
			t.Errorf("Expected false when getting non-existing key %d, got %v", key, value)
		}
	}

	// Test Has for non-existing keys
	for _, key := range nonExistingKeys {
		if sl.Has(key) {
			t.Errorf("Expected false for Has(%d)", key)
		}
	}
}

func TestSkipListGetMutable(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Set a value
	sl.Set(1, "original")

	// Get mutable reference and modify
	if ptr, exists := sl.GetMutable(1); exists {
		*ptr = "modified"
	} else {
		t.Error("Expected to get mutable reference for existing key")
	}

	// Verify the modification
	if value, exists := sl.Get(1); !exists || value != "modified" {
		t.Errorf("Expected 'modified', got %s", value)
	}

	// Test GetMutable for non-existing key
	if ptr, exists := sl.GetMutable(999); exists || ptr != nil {
		t.Error("Expected nil and false for non-existing key")
	}
}

func TestSkipListKeysValuesOrder(t *testing.T) {
	sl := NewOrderedSkipList[int, string]()

	// Insert keys in non-sorted order
	keys := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	values := []string{"three", "one", "four", "ONE", "five", "nine", "two", "six", "FIVE", "THREE"}

	for i, key := range keys {
		sl.Set(key, values[i])
	}

	// Get keys and verify they are sorted
	retrievedKeys := sl.Keys()
	expectedKeys := []int{1, 2, 3, 4, 5, 6, 9} // Sorted unique keys

	if !reflect.DeepEqual(retrievedKeys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, retrievedKeys)
	}

	// Get values and verify they are in order of their keys
	retrievedValues := sl.Values()
	expectedValues := []string{"ONE", "two", "THREE", "four", "FIVE", "six", "nine"} // Values for sorted keys

	if !reflect.DeepEqual(retrievedValues, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, retrievedValues)
	}

	// Get pairs and verify they are sorted by key
	retrievedPairs := sl.Pairs()
	expectedPairs := []pair.Pair[int, string]{
		{First: 1, Second: "ONE"},
		{First: 2, Second: "two"},
		{First: 3, Second: "THREE"},
		{First: 4, Second: "four"},
		{First: 5, Second: "FIVE"},
		{First: 6, Second: "six"},
		{First: 9, Second: "nine"},
	}

	if !reflect.DeepEqual(retrievedPairs, expectedPairs) {
		t.Errorf("Expected pairs %v, got %v", expectedPairs, retrievedPairs)
	}
}

func TestSkipListStringKeys(t *testing.T) {
	sl := NewOrderedSkipList[string, int]()

	// Insert test data with string keys
	testData := map[string]int{
		"banana": 2,
		"apple":  1,
		"cherry": 3,
		"date":   4,
	}

	for key, value := range testData {
		sl.Set(key, value)
	}

	// Test that keys are sorted lexicographically
	keys := sl.Keys()
	expectedKeys := []string{"apple", "banana", "cherry", "date"}

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}

	// Test values are in correct order
	values := sl.Values()
	expectedValues := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}
}