package huffman

import (
	"reflect"
	"testing"
)

func TestHuffman(t *testing.T) {
	testCases := []struct {
		in   []float64
		want []string
	}{{
		in:   []float64{0.1, 0.2, 0.45, 0.31, 0.9, 0.32},
		want: []string{"1000", "1001", "111", "101", "0", "110"},
	}, {
		in:   []float64{0.23, 0.24, 0.25, 0.26},
		want: []string{"00", "01", "10", "11"},
	}, {
		in:   []float64{0.5, 0.24, 0.25},
		want: []string{"1", "00", "01"},
	}, {
		in:   []float64{0.5, 0.49},
		want: []string{"1", "0"},
	}, {
		in:   []float64{1},
		want: []string{"0"},
	}, {
		in:   nil,
		want: nil,
	}}

	for _, tc := range testCases {
		if got := Huffman(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Huffman(%v) => got %v, want %v", tc.in, got, tc.want)
		}
	}
}
