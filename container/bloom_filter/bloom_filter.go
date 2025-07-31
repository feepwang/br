//go:build !go1.23
// +build !go1.23

package bloom_filter

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"math"
)

const (
	// defaultCapacity is the default expected number of items
	defaultCapacity = 1000
	// defaultFalsePositiveRate is the default acceptable false positive rate
	defaultFalsePositiveRate = 0.01 // 1%
)

// BloomFilter is a concrete implementation of the Interface.
type BloomFilter[T comparable] struct {
	bitArray          []bool  // The bit array storing the filter data
	bitSize           int     // Size of the bit array
	hashCount         int     // Number of hash functions
	capacity          int     // Expected number of items
	falsePositiveRate float64 // Target false positive rate
	itemCount         int     // Approximate number of items added
}

// NewBloomFilter creates a new Bloom filter with the specified capacity and false positive rate.
// If capacity is 0 or negative, defaultCapacity is used.
// If falsePositiveRate is 0 or negative, defaultFalsePositiveRate is used.
func NewBloomFilter[T comparable](capacity int, falsePositiveRate float64) Interface[T] {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		falsePositiveRate = defaultFalsePositiveRate
	}

	// Calculate optimal bit array size: m = -(n * ln(p)) / (ln(2)^2)
	// where n = capacity, p = false positive rate
	bitSize := int(math.Ceil(-float64(capacity) * math.Log(falsePositiveRate) / (math.Log(2) * math.Log(2))))

	// Calculate optimal number of hash functions: k = (m / n) * ln(2)
	hashCount := int(math.Ceil((float64(bitSize) / float64(capacity)) * math.Log(2)))

	// Ensure at least one hash function
	if hashCount < 1 {
		hashCount = 1
	}

	return &BloomFilter[T]{
		bitArray:          make([]bool, bitSize),
		bitSize:           bitSize,
		hashCount:         hashCount,
		capacity:          capacity,
		falsePositiveRate: falsePositiveRate,
		itemCount:         0,
	}
}

// NewBloomFilterWithDefaults creates a new Bloom filter with default parameters.
func NewBloomFilterWithDefaults[T comparable]() Interface[T] {
	return NewBloomFilter[T](defaultCapacity, defaultFalsePositiveRate)
}

// hash generates hash values for the given item using different seeds.
func (bf *BloomFilter[T]) hash(item T, seed uint32) uint32 {
	h := fnv.New32a()

	// Convert the item to bytes using a more robust method
	itemStr := fmt.Sprintf("%v", item)
	h.Write([]byte(itemStr))

	// Add seed bytes for variation
	seedBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(seedBytes, seed)
	h.Write(seedBytes)

	return h.Sum32()
}

// getHashIndices returns the hash indices for the given item.
func (bf *BloomFilter[T]) getHashIndices(item T) []int {
	indices := make([]int, bf.hashCount)
	for i := 0; i < bf.hashCount; i++ {
		hash := bf.hash(item, uint32(i))
		indices[i] = int(hash % uint32(bf.bitSize))
	}
	return indices
}

// Add inserts an item into the Bloom filter.
func (bf *BloomFilter[T]) Add(item T) {
	indices := bf.getHashIndices(item)
	for _, index := range indices {
		if !bf.bitArray[index] {
			bf.bitArray[index] = true
		}
	}
	bf.itemCount++
}

// Contains tests whether an item might be in the set.
func (bf *BloomFilter[T]) Contains(item T) bool {
	indices := bf.getHashIndices(item)
	for _, index := range indices {
		if !bf.bitArray[index] {
			return false
		}
	}
	return true
}

// Clear resets the Bloom filter to its initial empty state.
func (bf *BloomFilter[T]) Clear() {
	for i := range bf.bitArray {
		bf.bitArray[i] = false
	}
	bf.itemCount = 0
}

// Len returns the approximate number of items that have been added.
func (bf *BloomFilter[T]) Len() int {
	return bf.itemCount
}

// Capacity returns the estimated maximum number of items.
func (bf *BloomFilter[T]) Capacity() int {
	return bf.capacity
}

// FalsePositiveRate returns the current estimated false positive rate.
func (bf *BloomFilter[T]) FalsePositiveRate() float64 {
	if bf.itemCount == 0 {
		return 0.0
	}

	// Calculate current false positive rate: (1 - e^(-k*n/m))^k
	// where k = hash count, n = items added, m = bit array size
	exponent := -float64(bf.hashCount*bf.itemCount) / float64(bf.bitSize)
	base := 1.0 - math.Exp(exponent)
	return math.Pow(base, float64(bf.hashCount))
}

// BitSize returns the size of the underlying bit array.
func (bf *BloomFilter[T]) BitSize() int {
	return bf.bitSize
}

// HashCount returns the number of hash functions used.
func (bf *BloomFilter[T]) HashCount() int {
	return bf.hashCount
}
