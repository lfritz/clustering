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
