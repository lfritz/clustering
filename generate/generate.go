// Package generate implements functions that generate random distributions of points. Each function
// adds noise with a 2-D Normal (Gaussian) distribution.
package generate

import (
	"math"
	"math/rand"
)

// Point generates points around a given point.
func Point(n int, p [2]float64, stdDev float64) [][2]float64 {
	output := make([][2]float64, n)
	for i := range output {
		output[i][0] = p[0] + rand.NormFloat64()*stdDev
		output[i][1] = p[1] + rand.NormFloat64()*stdDev
	}
	return output
}

// CircularArc generates points along a circular arc, i.e. part of a circle. The center and radius
// arguments specify the circle; from and to are angles in radians that specify what part of the
// circle to include.
func CircularArc(n int, center [2]float64, radius, from, to float64, stdDev float64) [][2]float64 {
	output := make([][2]float64, n)
	d := to - from
	for i := range output {
		theta := from + d*rand.Float64()
		output[i][0] = center[0] + radius*math.Sin(theta) + rand.NormFloat64()*stdDev
		output[i][1] = center[1] + radius*math.Cos(theta) + rand.NormFloat64()*stdDev
	}
	return output
}
