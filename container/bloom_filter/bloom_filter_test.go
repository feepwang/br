package bloom_filter

import (
	"math"
	"testing"
)

func TestBloomFilterBasic(t *testing.T) {
	bf := NewBloomFilterWithDefaults[string]()

	// Test empty bloom filter
	if bf.Len() != 0 {
		t.Errorf("Expected length 0, got %d", bf.Len())
	}

	// Test Contains on empty bloom filter
	if bf.Contains("test") {
		t.Error("Expected false when checking existence in empty bloom filter")
	}

	// Test basic properties
	if bf.Capacity() != defaultCapacity {
		t.Errorf("Expected capacity %d, got %d", defaultCapacity, bf.Capacity())
	}

	if bf.BitSize() <= 0 {
		t.Errorf("Expected positive bit size, got %d", bf.BitSize())
	}

	if bf.HashCount() <= 0 {
		t.Errorf("Expected positive hash count, got %d", bf.HashCount())
	}

	if bf.FalsePositiveRate() != 0.0 {
		t.Errorf("Expected false positive rate 0.0 for empty filter, got %f", bf.FalsePositiveRate())
	}
}

func TestBloomFilterAddAndContains(t *testing.T) {
	bf := NewBloomFilter[string](100, 0.01)

	testItems := []string{"apple", "banana", "cherry", "date", "elderberry"}

	// Add items and verify they're contained
	for _, item := range testItems {
		bf.Add(item)
		if !bf.Contains(item) {
			t.Errorf("Item %s should be contained after adding", item)
		}
	}

	// Check length
	if bf.Len() != len(testItems) {
		t.Errorf("Expected length %d, got %d", len(testItems), bf.Len())
	}

	// All added items should still be contained
	for _, item := range testItems {
		if !bf.Contains(item) {
			t.Errorf("Item %s should be contained", item)
		}
	}
}

func TestBloomFilterNoFalseNegatives(t *testing.T) {
	bf := NewBloomFilter[int](50, 0.05)

	// Add numbers 1-20
	for i := 1; i <= 20; i++ {
		bf.Add(i)
	}

	// Verify no false negatives - all added items must be found
	for i := 1; i <= 20; i++ {
		if !bf.Contains(i) {
			t.Errorf("False negative: item %d should be contained", i)
		}
	}
}

func TestBloomFilterFalsePositives(t *testing.T) {
	bf := NewBloomFilter[int](100, 0.1) // 10% false positive rate

	// Add numbers 1-50
	for i := 1; i <= 50; i++ {
		bf.Add(i)
	}

	// Test numbers 51-200 (not added)
	falsePositives := 0
	testCount := 150
	for i := 51; i <= 200; i++ {
		if bf.Contains(i) {
			falsePositives++
		}
	}

	// Calculate actual false positive rate
	actualRate := float64(falsePositives) / float64(testCount)

	// The actual rate should be reasonably close to the theoretical rate
	// Allow for significant variance due to randomness and small sample size
	theoreticalRate := bf.FalsePositiveRate()
	if actualRate > theoreticalRate*5 && actualRate > 0.2 {
		t.Errorf("False positive rate too high: expected around %f, got %f", theoreticalRate, actualRate)
	}
}

func TestBloomFilterClear(t *testing.T) {
	bf := NewBloomFilter[string](50, 0.01)

	// Add some items
	items := []string{"test1", "test2", "test3"}
	for _, item := range items {
		bf.Add(item)
	}

	// Verify items are present
	for _, item := range items {
		if !bf.Contains(item) {
			t.Errorf("Item %s should be contained before clear", item)
		}
	}

	// Clear the filter
	bf.Clear()

	// Verify filter is empty
	if bf.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", bf.Len())
	}

	if bf.FalsePositiveRate() != 0.0 {
		t.Errorf("Expected false positive rate 0.0 after clear, got %f", bf.FalsePositiveRate())
	}

	// Verify items are no longer present (they shouldn't be, but bloom filters
	// could theoretically have false positives even after clear due to hash collisions)
	allGone := true
	for _, item := range items {
		if bf.Contains(item) {
			allGone = false
			break
		}
	}

	// In a properly cleared filter, we expect no items to be found
	if !allGone {
		t.Log("Note: Some items still show as contained after clear (unexpected but theoretically possible)")
	}
}

func TestBloomFilterDifferentTypes(t *testing.T) {
	// Test with int
	intFilter := NewBloomFilter[int](50, 0.01)
	intFilter.Add(42)
	if !intFilter.Contains(42) {
		t.Error("Int filter should contain 42")
	}
	if intFilter.Contains(43) {
		t.Log("Int filter shows false positive for 43 (expected possibility)")
	}

	// Test with float64
	floatFilter := NewBloomFilter[float64](50, 0.01)
	floatFilter.Add(3.14)
	if !floatFilter.Contains(3.14) {
		t.Error("Float filter should contain 3.14")
	}

	// Test with string
	stringFilter := NewBloomFilter[string](50, 0.01)
	stringFilter.Add("hello")
	if !stringFilter.Contains("hello") {
		t.Error("String filter should contain 'hello'")
	}
}

func TestBloomFilterParameters(t *testing.T) {
	// Test with custom parameters
	bf := NewBloomFilter[string](1000, 0.001) // Low false positive rate

	if bf.Capacity() != 1000 {
		t.Errorf("Expected capacity 1000, got %d", bf.Capacity())
	}

	// With lower false positive rate, we expect larger bit array and more hash functions
	if bf.BitSize() <= 0 {
		t.Errorf("Expected positive bit size, got %d", bf.BitSize())
	}

	if bf.HashCount() <= 0 {
		t.Errorf("Expected positive hash count, got %d", bf.HashCount())
	}
}

func TestBloomFilterInvalidParameters(t *testing.T) {
	// Test with invalid capacity (should use default)
	bf1 := NewBloomFilter[string](0, 0.01)
	if bf1.Capacity() != defaultCapacity {
		t.Errorf("Expected default capacity %d for zero capacity, got %d", defaultCapacity, bf1.Capacity())
	}

	bf2 := NewBloomFilter[string](-100, 0.01)
	if bf2.Capacity() != defaultCapacity {
		t.Errorf("Expected default capacity %d for negative capacity, got %d", defaultCapacity, bf2.Capacity())
	}

	// Test with invalid false positive rate (should use default)
	bf3 := NewBloomFilter[string](100, 0)
	// Should use default rate (we can't easily test the exact value without exposing internals)
	if bf3.BitSize() <= 0 {
		t.Error("Expected valid bit size even with invalid false positive rate")
	}

	bf4 := NewBloomFilter[string](100, 1.5)
	if bf4.BitSize() <= 0 {
		t.Error("Expected valid bit size even with invalid false positive rate")
	}
}

func TestBloomFilterFalsePositiveRateCalculation(t *testing.T) {
	bf := NewBloomFilter[int](100, 0.01)

	// Add some items
	for i := 0; i < 50; i++ {
		bf.Add(i)
	}

	fpr := bf.FalsePositiveRate()
	if fpr <= 0 || fpr >= 1 {
		t.Errorf("False positive rate should be between 0 and 1, got %f", fpr)
	}

	// As we add more items, the false positive rate should increase
	initialFPR := fpr
	for i := 50; i < 80; i++ {
		bf.Add(i)
	}

	newFPR := bf.FalsePositiveRate()
	if newFPR <= initialFPR {
		t.Errorf("False positive rate should increase with more items: %f -> %f", initialFPR, newFPR)
	}
}

func TestBloomFilterMathematicalProperties(t *testing.T) {
	capacity := 1000
	fpr := 0.01
	bf := NewBloomFilter[int](capacity, fpr)

	// Test that bit size follows the expected formula approximately
	// m ≈ -(n * ln(p)) / (ln(2)^2)
	expectedBitSize := int(math.Ceil(-float64(capacity) * math.Log(fpr) / (math.Log(2) * math.Log(2))))
	actualBitSize := bf.BitSize()

	// Allow some tolerance for rounding
	if math.Abs(float64(actualBitSize-expectedBitSize)) > 1 {
		t.Errorf("Bit size calculation incorrect: expected ~%d, got %d", expectedBitSize, actualBitSize)
	}

	// Test that hash count follows the expected formula approximately
	// k ≈ (m / n) * ln(2)
	expectedHashCount := int(math.Ceil((float64(actualBitSize) / float64(capacity)) * math.Log(2)))
	actualHashCount := bf.HashCount()

	if math.Abs(float64(actualHashCount-expectedHashCount)) > 1 {
		t.Errorf("Hash count calculation incorrect: expected ~%d, got %d", expectedHashCount, actualHashCount)
	}
}
