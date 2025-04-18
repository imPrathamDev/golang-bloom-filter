package bloombool

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash"
)

type BloomFilter struct {
	filter []bool
	size int
}

var (
	hasher = sha1.New()
)

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size: size,
	}
}

func (f *BloomFilter) Set(s string) {
	f.filter[f.hashPosition(s)] = true
}

func (f *BloomFilter) Check(s string) bool {
	return f.filter[f.hashPosition(s)]
}

// for creating hash
func createHash(h hash.Hash, input string) int {
	bits := h.Sum([]byte(input))
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	fmt.Println("result", result)
	return int(result) // cast the int64
}

// for getting position means index
func (f *BloomFilter) hashPosition(s string) int {
	hs := createHash(hasher, s)
	if hs < 0 {
		hs = -hs
	}
	return hs % len(f.filter)
}