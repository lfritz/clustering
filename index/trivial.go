package index

import (
	"github.com/lfritz/clustering/geometry"
)

// A TrivialIndex is an implementation of Index that doesn't actually speed up
// any operations.
type TrivialIndex struct {
	points [][2]float64
}

// NewTrivialIndex returns a new TrivialIndex.
func NewTrivialIndex(points [][2]float64) *TrivialIndex {
	return &TrivialIndex{points}
}

// Points returns the slice of points.
func (i *TrivialIndex) Points() [][2]float64 {
	return i.points
}

// BoundingBox returns the indices of all points within the given
// axis-aligned bounding box.
func (i *TrivialIndex) BoundingBox(bb *geometry.BoundingBox) []int {
	result := []int{}
	for i, p := range i.points {
		if bb.Contains(p) {
			result = append(result, i)
		}
	}
	return result
}
