//go:build go1.23
// +build go1.23

package bloom_filter

import (
	"math"
	"testing"
)

func TestBloomFilter123Basic(t *testing.T) {
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

func TestBloomFilter123AddAndContains(t *testing.T) {
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

func TestBloomFilter123NoFalseNegatives(t *testing.T) {
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

func TestBloomFilter123FalsePositives(t *testing.T) {
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
	theoreticalRate := bf.FalsePositiveRate()
	if actualRate > theoreticalRate*5 && actualRate > 0.2 {
		t.Errorf("False positive rate too high: expected around %f, got %f", theoreticalRate, actualRate)
	}
}

func TestBloomFilter123Clear(t *testing.T) {
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
}

func TestBloomFilter123DifferentTypes(t *testing.T) {
	// Test with int
	intFilter := NewBloomFilter[int](50, 0.01)
	intFilter.Add(42)
	if !intFilter.Contains(42) {
		t.Error("Int filter should contain 42")
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

func TestBloomFilter123Parameters(t *testing.T) {
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

func TestBloomFilter123MathematicalProperties(t *testing.T) {
	capacity := 1000
	fpr := 0.01
	bf := NewBloomFilter[int](capacity, fpr)

	// Test that bit size follows the expected formula approximately
	expectedBitSize := int(math.Ceil(-float64(capacity) * math.Log(fpr) / (math.Log(2) * math.Log(2))))
	actualBitSize := bf.BitSize()

	// Allow some tolerance for rounding
	if math.Abs(float64(actualBitSize-expectedBitSize)) > 1 {
		t.Errorf("Bit size calculation incorrect: expected ~%d, got %d", expectedBitSize, actualBitSize)
	}

	// Test that hash count follows the expected formula approximately
	expectedHashCount := int(math.Ceil((float64(actualBitSize) / float64(capacity)) * math.Log(2)))
	actualHashCount := bf.HashCount()

	if math.Abs(float64(actualHashCount-expectedHashCount)) > 1 {
		t.Errorf("Hash count calculation incorrect: expected ~%d, got %d", expectedHashCount, actualHashCount)
	}
}

func TestBloomFilter123PerformanceCharacteristics(t *testing.T) {
	// Test with a larger dataset to verify performance characteristics
	bf := NewBloomFilter[int](10000, 0.01)

	// Add a significant number of items
	itemCount := 5000
	for i := 0; i < itemCount; i++ {
		bf.Add(i)
	}

	// Verify all added items are found (no false negatives)
	for i := 0; i < itemCount; i++ {
		if !bf.Contains(i) {
			t.Errorf("False negative for item %d", i)
		}
	}

	// Test false positive rate with items not added
	falsePositives := 0
	testRange := 1000
	for i := itemCount; i < itemCount+testRange; i++ {
		if bf.Contains(i) {
			falsePositives++
		}
	}

	actualFPR := float64(falsePositives) / float64(testRange)
	theoreticalFPR := bf.FalsePositiveRate()

	// The actual FPR should be reasonably close to theoretical
	// Allow for more variance with smaller datasets
	if actualFPR > theoreticalFPR*10 && actualFPR > 0.1 {
		t.Errorf("Actual false positive rate %f significantly higher than theoretical %f", actualFPR, theoreticalFPR)
	}
}
