// Package random implements functions that generate random distributions of points.
package random

import (
	"math/rand"
)

// ChooseWeighted returns a pseudo-random int in 0..len(weights), chosen with probability
// proportional to the given weights. The elements of weights must be non-negative numbers and its
// sum must be positive.
func ChooseWeighted(weights []float64) int {
	return chooseWeighted(weights, rand.Float64())
}

// chooseWeighted implements the non-random part of ChooseWeighted, so it can be unit-tested.
func chooseWeighted(weights []float64, r float64) int {
	sum := 0.0
	for _, w := range weights {
		sum += w
	}
	r *= sum

	sum = 0
	for i, w := range weights {
		sum += w
		if r < sum {
			return i
		}
	}
	return len(weights) - 1
}

// Norm2D generates points with a 2-D normal (Gaussian) distribution.
func Norm2D(mean [2]float64, stdDev float64, output [][2]float64) {
	for i := range output {
		x := rand.NormFloat64()*stdDev + mean[0]
		y := rand.NormFloat64()*stdDev + mean[1]
		output[i][0] = x
		output[i][1] = y
	}
}
