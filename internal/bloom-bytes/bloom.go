package bloombytes

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

/*
* By using Byte we can save space
 */

type Bloom struct {
	filter []byte
	size   int
	hashFuncs int
}

const SEED = 52342;

func NewBloom(size int, hashFuncs int) *Bloom {
	numBytes := (size + 7) / 8
	return &Bloom{
		filter: make([]byte, numBytes),
		size:   size,
		hashFuncs: hashFuncs,
	}
}

func (b *Bloom) setBit(p uint) {
	pos := p / 8
	bitPos := p % 8
	b.filter[pos] = b.filter[pos] | 1 << bitPos
}

func (b *Bloom) getBit(p uint) byte {
	pos, bitPos := p/8, p%8
	return b.filter[pos] & (1 << bitPos)
}

func (b *Bloom) hash(key string, seed int) uint {
	h := murmur3.New64WithSeed(uint32(seed))
	h.Write([]byte(key))
	hashValue := h.Sum64()
fmt.Println("len", uint(hashValue) % uint(b.size))
	return uint(hashValue) % uint(b.size)
}

func (b *Bloom) Set(key string) {
	for i := 0; i < b.hashFuncs; i++ {
		position := b.hash(key, i)
		b.setBit(position)
	}
}

func (b *Bloom) Check(key string) bool {
	for i := 0; i < b.hashFuncs; i++ {
		position := b.hash(key, i)
		a := b.getBit(position)
		return a != 0
	}

	return true
}