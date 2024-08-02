package bloomfilter

import (
	"fmt"
	"math/rand"

	"github.com/nithishravindra/bloom-filter/internal/hashing"
)

// BloomFilter represents a Bloom filter data structure
type BloomFilter struct {
	Filter []uint8
	Size   int32
}

// NewBloomFilter creates a new Bloom filter with a specified size
func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		Filter: make([]uint8, size),
		Size:   size,
	}
}

// Add adds a key to the Bloom filter using a specified number of hash functions
func (b *BloomFilter) Add(key string, numHashFn int) {
	for i := 0; i < numHashFn; i++ {
		hasher := hashing.NewMurmurHasher(uint32(rand.Uint32())) // Create a new hasher with a random seed
		idx := hasher.Hash(key)
		aIdx := idx / 8
		bIdx := idx % 8
		b.Filter[aIdx] |= (1 << bIdx)
	}
}

// Exists checks if a key might exist in the Bloom filter using a specified number of hash functions
func (b *BloomFilter) Exists(key string, numHashFn int) bool {
	for i := 0; i < numHashFn; i++ {
		hasher := hashing.NewMurmurHasher(uint32(rand.Uint32())) // Create a new hasher with a random seed
		idx := hasher.Hash(key)
		aIdx := idx / 8
		bIdx := idx % 8
		if b.Filter[aIdx]&(1<<bIdx) == 0 {
			return false
		}
	}
	return true
}

// Print prints the Bloom filter's contents (for debugging purposes)
func (b *BloomFilter) Print() {
	fmt.Println(b.Filter)
}
