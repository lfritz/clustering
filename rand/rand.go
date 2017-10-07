// Package rand implements functions that generate random distributions of points.
package rand

import (
	"math/rand"
)

// Norm2D generates points with a 2-D normal (Gaussian) distribution.
func Norm2D(mean [2]float64, stdDev float64, output [][2]float64) {
	for i := range output {
		x := rand.NormFloat64()*stdDev + mean[0]
		y := rand.NormFloat64()*stdDev + mean[1]
		output[i][0] = x
		output[i][1] = y
	}
}
