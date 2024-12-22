package iters

import (
	"testing"

	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

func TestFloats(t *testing.T) {
	tests := []struct {
		name  string
		iter  *sFloats.Iter
		count int
	}{
		{"v1", sFloats.NewIter(0, 10, 1.1, 2, 12.3, 45, 6.6), 7},
		{"v2", sFloats.NewIter(3, 3, 1.2, 3, 12, 45.7, 7.2), 7},
		{"v3", sFloats.NewIter(3, 0, 1, 1, 12, 45, 6), 6},
		{"v4", sFloats.NewIter(0, 0, 0.1, 2, 12, 45, 6), 6},
		{"v5", sFloats.NewIter(0, 10, 1, 1), 1},
		{"v6", sFloats.NewIter(0, 10, 0, 2), 2},
		{"v7", sFloats.NewIter(0, 0, 0, 1), 1},
		{"v8", sFloats.NewIter(0, 0, 0, 1, 0), 1},
		{"v9", sFloats.NewIter(10, 0, 1, 2), 2},
		{"v10", sFloats.NewIter(10, 0, -0.15, 3), 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			for tt.iter.Next() {
				count++
			}
			if count != tt.count {
				t.Errorf("Expected %d, got %d", tt.count, count)
			}
		})
	}
}

func TestInts(t *testing.T) {
	tests := []struct {
		name  string
		iter  *sInts.Iter
		count int
	}{
		{"v1", sInts.NewIter(0, 10, 1, 12, 45, 6), 6},
		{"v2", sInts.NewIter(3, 3, 1, 12, 45, 6), 6},
		{"v3", sInts.NewIter(3, 0, 1, 12, 45, 6), 6},
		{"v4", sInts.NewIter(0, 0, 1, 12, 45, 6), 6},
		{"v5", sInts.NewIter(0, 10, 1), 1},
		{"v6", sInts.NewIter(0, 10, 0), 1},
		{"v7", sInts.NewIter(0, 0, 0), 1},
		{"v8", sInts.NewIter(0, 0, 0, 0), 1},
		{"v9", sInts.NewIter(2, 2, 0), 1},
		{"v10", sInts.NewIter(2, 2, 0, 2), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			for tt.iter.Next() {
				count++
			}
			if count != tt.count {
				t.Errorf("Expected %d, got %d", tt.count, count)
			}
		})
	}
}
