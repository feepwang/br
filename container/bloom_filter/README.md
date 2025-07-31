# Bloom Filter

A Bloom Filter is a space-efficient probabilistic data structure designed to test whether an element is a member of a set. False positive matches are possible, but false negatives are not.

## Features

- **Generic Type Support**: Works with any comparable type (int, string, float64, etc.)
- **Configurable Parameters**: Set expected capacity and false positive rate
- **Optimal Performance**: Uses mathematically optimal bit array size and hash function count
- **Go Version Support**: Compatible with Go 1.21+ and Go 1.23+ with separate implementations

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/feepwang/br/container/bloom_filter"
)

func main() {
    // Create a Bloom filter with capacity for 1000 items and 1% false positive rate
    bf := bloom_filter.NewBloomFilter[string](1000, 0.01)
    
    // Add items
    bf.Add("apple")
    bf.Add("banana")
    bf.Add("cherry")
    
    // Test membership
    fmt.Println(bf.Contains("apple"))  // true (definitely in set)
    fmt.Println(bf.Contains("grape"))  // false (definitely not in set)
    fmt.Println(bf.Contains("date"))   // might be true (possible false positive)
    
    // Get statistics
    fmt.Printf("Items added: %d\n", bf.Len())
    fmt.Printf("Capacity: %d\n", bf.Capacity())
    fmt.Printf("False positive rate: %.4f\n", bf.FalsePositiveRate())
    fmt.Printf("Bit array size: %d\n", bf.BitSize())
    fmt.Printf("Hash functions: %d\n", bf.HashCount())
}
```

### Default Parameters

```go
// Create with default parameters (1000 capacity, 1% false positive rate)
bf := bloom_filter.NewBloomFilterWithDefaults[string]()
```

### Different Types

```go
// Integer Bloom filter
intFilter := bloom_filter.NewBloomFilter[int](500, 0.05)
intFilter.Add(42)
fmt.Println(intFilter.Contains(42)) // true

// Float Bloom filter  
floatFilter := bloom_filter.NewBloomFilter[float64](200, 0.02)
floatFilter.Add(3.14159)
fmt.Println(floatFilter.Contains(3.14159)) // true
```

## Interface

The Bloom filter implements the following interface:

```go
type Interface[T comparable] interface {
    Add(item T)                    // Add an item to the filter
    Contains(item T) bool          // Test if item might be in the set
    Clear()                        // Reset the filter
    Len() int                      // Approximate number of items added
    Capacity() int                 // Maximum recommended items
    FalsePositiveRate() float64    // Current estimated false positive rate
    BitSize() int                  // Size of underlying bit array
    HashCount() int                // Number of hash functions used
}
```

## Mathematical Background

The Bloom filter uses optimal parameters calculated from:

- **Bit array size**: `m = -(n * ln(p)) / (ln(2)²)`
- **Hash functions**: `k = (m / n) * ln(2)`
- **False positive rate**: `(1 - e^(-k*n/m))^k`

Where:
- `n` = expected number of items
- `p` = desired false positive rate  
- `m` = bit array size
- `k` = number of hash functions

## Performance

- **Time Complexity**: O(k) for both Add and Contains operations, where k is the number of hash functions
- **Space Complexity**: O(m) where m is the bit array size
- **No False Negatives**: If Contains returns false, the item is definitely not in the set
- **Possible False Positives**: If Contains returns true, the item might be in the set