package hashing

import (
	"hash"

	"github.com/spaolacci/murmur3"
)

// MurmurHasher provides a hasher object with seed capabilities
type MurmurHasher struct {
	hash.Hash32
	seed uint32
}

// NewMurmurHasher creates a new MurmurHasher instance with a seed
func NewMurmurHasher(seed uint32) *MurmurHasher {
	return &MurmurHasher{
		Hash32: murmur3.New32WithSeed(seed),
		seed:   seed,
	}
}

// Hash calculates the murmur hash for a given key
func (h *MurmurHasher) Hash(key string) int32 {
	h.Write([]byte(key))
	result := h.Sum32() % uint32(10000) // Assuming 10000 is the maximum size
	h.Reset()
	return int32(result)
}
