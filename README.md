# Bloom Filter Implementation

This Go project implements a Bloom filter data structure for efficiently checking if an element might be present in a set. It uses multiple MurmurHash3 hash functions for improved performance.

### packages

- hashing: Contains the MurmurHasher struct and associated methods for hashing strings.
- bloomfilter: Defines the BloomFilter struct and its methods for adding and checking elements.
- main: Contains the main function for testing the Bloom filter with randomly generated datasets.
