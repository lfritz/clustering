package index

import (
	"github.com/lfritz/clustering"
)

// A TrivialIndex is an implementation of Index that doesn't actually speed up
// any operations.
type TrivialIndex struct {
	points []clustering.Point
}

// NewTrivialIndex returns a new TrivialIndex.
func NewTrivialIndex(points []clustering.Point) *TrivialIndex {
	return &TrivialIndex{points}
}

func (i *TrivialIndex) Points() []clustering.Point {
	return i.points
}

func (i *TrivialIndex) BoundingBox(x0, x1, y0, y1 float64) []int {
	result := []int{}
	for i, p := range i.points {
		if x0 <= p.X && p.X < x1 && y0 <= p.Y && p.Y < y1 {
			result = append(result, i)
		}
	}
	return result
}

func (i *TrivialIndex) Circle(center clustering.Point, radius float64) []int {
	return circle(i, center, radius)
}
