//go:build go1.23
// +build go1.23

package ordered_map

import (
	"testing"
)

func TestRedBlackTreeIterators(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Set(3, "three")
	tree.Set(1, "one")
	tree.Set(2, "two")

	// Test KeySeq
	var keys []int
	for k := range tree.KeySeq() {
		keys = append(keys, k)
	}
	expectedKeys := []int{1, 2, 3}
	if len(keys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, got %d", len(expectedKeys), len(keys))
	}
	for i, key := range keys {
		if key != expectedKeys[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expectedKeys[i], key)
		}
	}

	// Test ValueSeq
	var values []string
	for v := range tree.ValueSeq() {
		values = append(values, v)
	}
	expectedValues := []string{"one", "two", "three"}
	if len(values) != len(expectedValues) {
		t.Errorf("Expected %d values, got %d", len(expectedValues), len(values))
	}
	for i, value := range values {
		if value != expectedValues[i] {
			t.Errorf("At index %d, expected %s, got %s", i, expectedValues[i], value)
		}
	}

	// Test PairSeq
	var pairs [][2]interface{}
	for k, v := range tree.PairSeq() {
		pairs = append(pairs, [2]interface{}{k, v})
	}
	expectedPairs := [][2]interface{}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
	}
	if len(pairs) != len(expectedPairs) {
		t.Errorf("Expected %d pairs, got %d", len(expectedPairs), len(pairs))
	}
	for i, pair := range pairs {
		if pair[0] != expectedPairs[i][0] || pair[1] != expectedPairs[i][1] {
			t.Errorf("At index %d, expected (%v, %v), got (%v, %v)",
				i, expectedPairs[i][0], expectedPairs[i][1], pair[0], pair[1])
		}
	}

	// Test early termination
	count := 0
	for k := range tree.KeySeq() {
		count++
		if k == 2 {
			break
		}
	}
	if count != 2 {
		t.Errorf("Expected to stop at 2 iterations, got %d", count)
	}
}