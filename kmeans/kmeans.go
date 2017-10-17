// Package kmeans implements the k-means clustering algorithm, with the k-means++ algorithm for
// choosing initial values.
//
// References:
//
//     https://en.m.wikipedia.org/wiki/K-means_clustering
//     https://en.m.wikipedia.org/wiki/K-means++
package kmeans

import (
	"github.com/lfritz/clustering/geometry"
	"math"
	"math/rand"
	"reflect"
)

// Repeat applies the k-means algorithm multiple times and returns the best result, i.e. the
// clustering with the lowest variance.
func Repeat(points [][2]float64, k int, repetitions int) ([]int, float64) {
	var bestCl []int
	lowestVariance := math.Inf(+1)
	for i := 0; i < repetitions; i++ {
		cl, variance := Kmeans(points, k)
		if variance < lowestVariance {
			bestCl = cl
			lowestVariance = variance
		}
	}
	return bestCl, lowestVariance
}

// Kmeans applies the k-means algorithm and returns a clustering that groups points into k clusters,
// along with the variance of the clustering.
func Kmeans(points [][2]float64, k int) ([]int, float64) {
	centroids := initPlusPlus(points, k)
	var cl []int
	var variance float64
	for {
		nextClustering, nextVariance := clusteringForCentroids(points, centroids)
		if reflect.DeepEqual(nextClustering, cl) {
			break
		}
		cl = nextClustering
		variance = nextVariance
		centroids = centroidsForClusters(points, k, cl)

	}
	return cl, variance
}

// randK randomly selects k numbers in [0..n), without duplicates.
func randK(k, n int) []int {
	result := make([]int, k)
	for i := range result {
		x := rand.Intn(n - i)
		done := false
		for j, other := range result[:i] {
			if other > x {
				copy(result[j+1:], result[j:])
				result[j] = x
				done = true
				break
			}
			x++
		}
		if !done {
			result[i] = x
		}
	}
	return result
}

// chooseWeighted returns a pseudo-random int in 0..len(weights), chosen with probability
// proportional to the given weights. The elements of weights must be non-negative numbers and its
// sum must be positive.
func chooseWeighted(weights []float64) int {
	return chooseWeightedFor(weights, rand.Float64())
}

// chooseWeightedFor implements the non-random part of chooseWeighted, so it can be unit-tested.
func chooseWeightedFor(weights []float64, r float64) int {
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

// initPlusPlus generates an initial set of centroids for the k-means algorithm following the
// k-means++ algorithm.
func initPlusPlus(points [][2]float64, k int) [][2]float64 {
	n := len(points)
	centroids := make([][2]float64, 0, n)

	// choose the first centroid randomly from points
	centroids = append(centroids, points[rand.Intn(n)])

	// for the remaining k-1 centroids, use a weighted probability distribution
	weights := make([]float64, len(points))
	for i := 1; i < k; i++ {
		// for each point, compute its squared distance from the closest centroid
		for j, p := range points {
			_, weights[j] = closest(centroids, p)
		}
		centroids = append(centroids, points[chooseWeighted(weights)])
	}

	return centroids
}

func closest(ps [][2]float64, q [2]float64) (int, float64) {
	closest := 0
	minimum := math.Inf(+1)
	for i, p := range ps {
		val := geometry.DistanceSquared(p, q)
		if val < minimum {
			minimum = val
			closest = i
		}
	}
	return closest, minimum
}

// clusteringForCentroids takes a set of points and a set of centroids and returns a clustering that
// assigns each point to the closest centroid, along with its variance.
func clusteringForCentroids(points [][2]float64, centroids [][2]float64) ([]int, float64) {
	cl := make([]int, len(points))
	variance := 0.0
	var v float64
	for i, p := range points {
		cl[i], v = closest(centroids, p)
		variance += v
	}
	return cl, variance
}

// centroidsForClusters calculates a centroid for each cluster as the mean of its points.
func centroidsForClusters(points [][2]float64, k int, cl []int) [][2]float64 {
	count := make([]int, k)
	centroids := make([][2]float64, k)
	for i, c := range cl {
		centroids[c][0] += points[i][0]
		centroids[c][1] += points[i][1]
		count[c]++
	}
	for i, c := range count {
		if c != 0 {
			centroids[i][0] /= float64(c)
			centroids[i][1] /= float64(c)
		}
	}
	return centroids
}
