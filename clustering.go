package clustering

import (
	"math"
)

// A Point is a point in 2-dimensional Euclidian space.
type Point struct {
	X, Y float64
}

// Equal returns true if q has the same value as p.
func (p Point) Equal(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

// Distance returns the Euclidian distance between two points.
func Distance(a, b Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// A spatial index for a slice of points.
type PointIndex interface {
	// Points returns the slice of points.
	Points() []Point
	// BoundingBox returns the indices of all points within the given
	// axis-aligned bounding box.
	BoundingBox(x0, x1, y0, y1 float64) []int
	// Neighbors returns the indices of all points in the circle with the
	// given center and radius.
	Neighbors(center Point, r float64) []int
}
