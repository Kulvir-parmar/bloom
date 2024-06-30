package bloomfilter

import (
	"errors"
	"fmt"
	"math"

	mmh "github.com/spaolacci/murmur3"
)

type bf struct {
	numHashFunctions int
	buffer           []uint64
}

func NewBloomFilter(fpRate float64, expectedInsertions int) (*bf, error) {
	if fpRate < 0.0 {
		return nil, errors.New("False Positive Rate must be greater than 0")
	}
	if fpRate >= 1.0 {
		return nil, errors.New("False Positive Rate must be smaller than 1")
	}
	if expectedInsertions <= 0 {
		return nil, errors.New("No of insertions must be greater than 0")
	}

	numBits := numBits(fpRate, expectedInsertions)
	numHashFunction := numHasFunctions(numBits, expectedInsertions)
	size := uint(math.Ceil(float64(numBits) / 64.0))

	return &bf{
		numHashFunctions: numHashFunction,
		buffer:           make([]uint64, size),
	}, nil
}

// num of bits required for given insertions (n) and fp rate (e) are given by formula
// numBits = Ceil((n * log(e)) / log(1 / pow(2, log(2))))
func numBits(fpRate float64, expectedInsertions int) int {
	return int(math.Ceil((float64(expectedInsertions) * math.Log(fpRate)) / math.Log(1/math.Pow(2, math.Log(2)))))
}

// num of hash functions recommended for given bits(m) and expectedInsertions(n) are given by formula
// numHashFunctions = round((m / n) * log(2))
func numHasFunctions(numBits, expectedInsertions int) int {
	return int(math.Round((float64(numBits) / float64(expectedInsertions)) * math.Log(2)))
}

// hashing optimizations
// https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf
// i(x) = h1(x) + ih2(x) mod m
func (bf *bf) Put(key interface{}) {
	bytes := GetBytes(key)
	if len(bytes) == 0 {
		fmt.Printf("Failed to convert %v to bytes array /n", key)
		return
	}

	hash64, _ := mmh.Sum128(bytes)
	hash1 := int(hash64)
	hash2 := int(hash64 >> 32)

	for num := range bf.numHashFunctions {
		combinedHash := hash1 + (num * hash2)
		if combinedHash < 0 {
			combinedHash = ^combinedHash
		}
		bf.setBit(combinedHash)
	}
}

func (bf *bf) MightContain(key interface{}) bool {
	bytes := GetBytes(key)
	if len(bytes) == 0 {
		fmt.Printf("Failed to convert %v to bytes array /n", key)
		return false
	}

	hash64, _ := mmh.Sum128(bytes)
	hash1 := int(hash64)
	hash2 := int(hash64 >> 32)

	for num := range bf.numHashFunctions {
		combinedHash := hash1 + (num * hash2)
		if combinedHash < 0 {
			combinedHash = ^combinedHash
		}

		if !bf.checkSetBit(combinedHash) {
			return false
		}
	}
	return true
}

func (bf *bf) checkSetBit(hash int) bool {
	idx := hash % (len(bf.buffer) * 64)
	byteIdx, bitIdx := getIdx(idx)

	if bf.buffer[byteIdx]&(1<<bitIdx) == 0 {
		return false
	}
	return true
}

func (bf *bf) setBit(hash int) {
	idx := hash % (len(bf.buffer) * 64)
	byteIdx, bitIdx := getIdx(idx)
	bf.buffer[byteIdx] |= (1 << bitIdx)
}

func getIdx(idx int) (int, int) {
	return idx / 64, idx % 64
}
