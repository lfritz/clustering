// Package rand implements functions that generate random distributions of points.
package rand

import (
	"github.com/lfritz/clustering"
	"math/rand"
)

// Norm2D generates points with a 2-D normal (Gaussian) distribution.
func Norm2D(mean clustering.Point, stdDev float64, output []clustering.Point) {
	for i := range output {
		x := rand.NormFloat64()*stdDev + mean.X
		y := rand.NormFloat64()*stdDev + mean.Y
		output[i] = clustering.Point{X: x, Y: y}
	}
}
