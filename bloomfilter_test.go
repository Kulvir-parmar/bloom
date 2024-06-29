package bloomfilter

import (
	"testing"
)

func TestNumBits(t *testing.T) {
	type tc struct {
		numInsertions int
		fpRate        float64
		numBits       int
	}

	cases := []tc{{10000, 0.001, 143776}, {1000000, 0.01, 9585059}, {1000000, 0.5, 1442696}}

	for _, c := range cases {
		bits := numBits(c.fpRate, c.numInsertions)

		if bits != c.numBits {
			t.Errorf("Expected %d and got %d number of bits", c.numBits, bits)
		}
	}
}

func TestNumHashFunctions(t *testing.T) {
	type tc struct {
		numBits       int
		numInsertions int
		numHashes     int
	}

	cases := []tc{{9585059, 1000000, 7}, {14378, 1000, 10}, {19171, 1000, 13}}

	for _, c := range cases {
		hashes := numHasFunctions(c.numBits, c.numInsertions)
		if hashes != c.numHashes {
			t.Errorf("Expected %d and got %d number of hash functions", c.numHashes, hashes)
		}
	}
}

func TestGetIdx(t *testing.T) {
}
