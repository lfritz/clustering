package index

import (
	"github.com/lfritz/clustering"
)

// A spatial index for a slice of points.
type Index interface {
	// Points returns the slice of points.
	Points() []clustering.Point
	// BoundingBox returns the indices of all points within the given
	// axis-aligned bounding box.
	BoundingBox(x0, x1, y0, y1 float64) []int
	// Circle returns the indices of all points in the circle with the
	// given center and radius.
	Circle(center clustering.Point, radius float64) []int
}

// circle implements Circle in terms of BoundingBox.
func circle(i Index, center clustering.Point, radius float64) []int {
	inBB := i.BoundingBox(center.X-radius, center.X+radius, center.Y-radius, center.Y+radius)
	points := i.Points()
	result := []int{}
	for _, p := range inBB {
		if clustering.Distance(center, points[p]) <= radius {
			result = append(result, p)
		}
	}
	return result
}
