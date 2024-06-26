package bloomfilter

import (
	"errors"
	"math"
)

type bf struct {
	numHashFunctions int
	buffer           []uint64
}

func NewBloomFilter(fpRate float64, expectedInsertions int64) (*bf, error) {
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
func numBits(fpRate float64, expectedInsertions int64) int64 {
	return int64(math.Ceil((float64(expectedInsertions) * math.Log(fpRate)) / math.Log(1/math.Pow(2, math.Log(2)))))
}

// num of hash functions recommended for given bits(m) and expectedInsertions(n) are given by formula
// numHashFunctions = round((m / n) * log(2))
func numHasFunctions(numBits, expectedInsertions int64) int {
	return int(math.Round((float64(numBits) / float64(expectedInsertions)) * math.Log(2)))
}
