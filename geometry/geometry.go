// Package geometry implement some basic geometry functions for 2-D points.
package geometry

import (
	"math"
)

// Distance returns the Euclidian distance between two points.
func Distance(a, b [2]float64) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return math.Sqrt(dx*dx + dy*dy)
}
