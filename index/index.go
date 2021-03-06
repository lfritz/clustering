// Package index provides spatial indices for 2-D points.
package index

import (
	"github.com/lfritz/clustering/geometry"
)

// An Index is a spatial index for a slice of points.
type Index interface {
	// Points returns the slice of points.
	Points() [][2]float64
	// BoundingBox returns the indices of all points within the given
	// axis-aligned bounding box.
	BoundingBox(bb *geometry.BoundingBox) []int
}

// Circle returns the indices of all points in the circle with the given center
// and radius.
func Circle(i Index, c [2]float64, r float64) []int {
	bb := geometry.BoundingBox{
		From: [2]float64{c[0] - r, c[1] - r},
		To:   [2]float64{c[0] + r, c[1] + r},
	}
	inBB := i.BoundingBox(&bb)
	points := i.Points()
	result := []int{}
	rSquared := r * r
	for _, p := range inBB {
		if geometry.DistanceSquared(c, points[p]) <= rSquared {
			result = append(result, p)
		}
	}
	return result
}
