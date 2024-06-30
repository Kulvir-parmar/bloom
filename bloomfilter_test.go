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
	type tc struct {
		idx     int
		byteIdx int
		bitIdx  int
	}

	cases := []tc{
		{64, 1, 0},
		{28, 0, 1},
		{1024, 15, 40},
	}

	for _, c := range cases {
		byteidx, bitIdx := getIdx(c.bitIdx)

		if byteidx != c.byteIdx && bitIdx != c.bitIdx {
			t.Errorf("Expected byte idx %d and bit dx % d, got byte idx %d and bitIdx %d", c.byteIdx, c.bitIdx, byteidx, bitIdx)
		}
	}
}

func TestInsertions(t *testing.T) {
	bf, err := NewBloomFilter(0.01, 420)
	if err != nil {
		t.Errorf("Unexpected error during initialization of bloom filter : %v", err)
	}

	for i := 69; i < 169; i++ {
		bf.Put(i)
	}

	for i := 69; i < 169; i++ {
		if !bf.MightContain(i) {
			t.Fatalf("Expected %d to be present in the filter!!!", i)
		}
	}
}

func TestFpRate(t *testing.T) {
	bf, err := NewBloomFilter(0.1, 200)
	if err != nil {
		t.Errorf("Unexpected error during initialization of bloom filter : %v", err)
	}

	for i := 0; i < 120; i++ {
		bf.Put(i)
	}

	var errors int
	for i := 0; i < 200; i++ {
		if i >= 120 {
			if bf.MightContain(i) {
				errors++
			}
		}
	}
	fp := float64(errors) / float64((200 - 119))

	if fp > 0.1 {
		t.Fatalf("False positive rate should be less than 0.1, got %f with total errors %d", fp, errors)
	}

	// new bloom filter to test same thing again
	newBf, err := NewBloomFilter(0.069, 2000)
	if err != nil {
		t.Errorf("Unexpected error during initialization of bloom filter : %v", err)
	}

	for i := 0; i < 1500; i++ {
		newBf.Put(i)
	}

	errors = 0
	for i := 0; i < 2000; i++ {
		if i >= 1500 {
			if newBf.MightContain(i) {
				errors++
			}
		}
	}
	newFp := float64(errors) / float64((2000 - 1499))

	if newFp > 0.069 {
		t.Fatalf("False positive rate should be less than 0.069, got %f with total errors %d", newFp, errors)
	}
}
