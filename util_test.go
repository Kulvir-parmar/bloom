package bloomfilter

import (
	"bytes"
	"testing"
)

func TestGetBytes(t *testing.T) {
	type tc struct {
		key      interface{}
		expected []byte
	}
	cases := []tc{
		{"filter", []byte{102, 105, 108, 116, 101, 114}},
		{12345, []byte{57, 48, 0, 0}},
		{int32(12345), []byte{57, 48, 0, 0}},
		{int32(-12345), []byte{199, 207, 255, 255}},
		{3147483647, []byte{255, 201, 154, 187, 0, 0, 0, 0}},
		{int64(3147483647), []byte{255, 201, 154, 187, 0, 0, 0, 0}},
		{int64(-3147483647), []byte{1, 54, 101, 68, 255, 255, 255, 255}},
		{uint64(123456), []byte{64, 226, 1, 0, 0, 0, 0, 0}},
		{uint32(123456), []byte{64, 226, 1, 0}},
		{4611686018427387905, []byte{1, 0, 0, 0, 0, 0, 0, 64}},
		{[]byte{1, 2, 3}, []byte{1, 2, 3}},
		{[]int{1, 2, 3}, []byte{}},
	}

	for _, c := range cases {
		b := GetBytes(c.key)
		if !bytes.Equal(b, c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, b)
		}
	}
}
