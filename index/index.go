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
	BoundingBox(x0, x1, y0, y1 float64) []int
	// Circle returns the indices of all points in the circle with the
	// given center and radius.
	Circle(center [2]float64, radius float64) []int
}

// circle implements Circle in terms of BoundingBox.
func circle(i Index, center [2]float64, radius float64) []int {
	inBB := i.BoundingBox(center[0]-radius, center[0]+radius, center[1]-radius, center[1]+radius)
	points := i.Points()
	result := []int{}
	for _, p := range inBB {
		if geometry.Distance(center, points[p]) <= radius {
			result = append(result, p)
		}
	}
	return result
}
