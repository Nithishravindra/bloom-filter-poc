package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/nithishravindra/bloom-filter/internal/bloomfilter"
)

func main() {
	dataset := make([]string, 0)
	dataset_exists := make(map[string]bool)
	dataset_notexists := make(map[string]bool)

	// Generate datasets
	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_exists[u.String()] = true
	}

	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_notexists[u.String()] = true
	}

	// Test Bloom filter with different number of hash functions
	for j := 1; j <= 100; j++ { // Adjust the maximum number of hash functions as needed
		bloom := bloomfilter.NewBloomFilter(10000) // Choose a fixed size for the Bloom filter

		log.Printf("Testing with %d hash functions\n", j)

		// Add existing elements to the Bloom filter
		for key := range dataset_exists {
			bloom.Add(key, j)
		}

		// Calculate false positive rate
		falsePositive := 0
		for _, key := range dataset {
			exists := bloom.Exists(key, j)
			if exists {
				if _, ok := dataset_notexists[key]; ok {
					falsePositive++
				}
			}
		}
		falsePositiveRate := float64(falsePositive) / float64(len(dataset_notexists))
		log.Printf("False positive rate: %.4f\n", falsePositiveRate)
	}
}
