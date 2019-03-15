package main

import "hash/fnv"

// Hash returns hashing number.
func Hash(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	sum := h.Sum32()
	return sum
}
