// Package kmeans implements the k-means clustering algorithm.
//
// References:
//
//     https://en.m.wikipedia.org/wiki/K-means_clustering
package kmeans

import (
	"github.com/lfritz/clustering/geometry"
	"github.com/lfritz/clustering/rand"
	"math"
	stdRand "math/rand"
	"reflect"
)

// Kmeans applies the k-means algorithm and returns a clustering that groups points into k clusters.
// It uses the Forgy method to initialize the clusters, i.e. it randomly chooses k observations from
// the data set.
func Kmeans(points [][2]float64, k int) []int {
	centroids := initPlusPlus(points, k)
	var cl []int
	for {
		nextClustering := clusteringForCentroids(points, centroids)
		if reflect.DeepEqual(nextClustering, cl) {
			break
		}
		cl = nextClustering
		centroids = centroidsForClusters(points, k, cl)

	}
	return cl
}

// TODO:
// - compare the two initialization methods
// - try running k-means repeatedly

// randK randomly selects k numbers in [0..n), without duplicates.
func randK(k, n int) []int {
	result := make([]int, k)
	for i := range result {
		x := stdRand.Intn(n - i)
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

// initForgy generates an initial set of centroids for the k-means algorithm using the Forgy method.
func initForgy(points [][2]float64, k int) [][2]float64 {
	centroids := make([][2]float64, k)
	for i, x := range randK(k, len(points)) {
		centroids[i] = points[x]
	}
	return centroids
}

// initPlusPlus generates an initial set of centroids for the k-means algorithm following the
// k-means++ algorithm.
func initPlusPlus(points [][2]float64, k int) [][2]float64 {
	n := len(points)
	centroids := make([][2]float64, 0, n)

	// choose the first centroid randomly from points
	centroids = append(centroids, points[stdRand.Intn(n)])

	// for the remaining k-1 centroids, use a weighted probability distribution
	weights := make([]float64, len(points))
	for i := 1; i < k; i++ {
		// for each point, compute its squared distance from the closest centroid
		for j, p := range points {
			weights[j] = geometry.DistanceSquared(p, centroids[closest(centroids, p)])
		}
		centroids = append(centroids, points[rand.ChooseWeighted(weights)])
	}

	return centroids
}

func closest(ps [][2]float64, q [2]float64) int {
	closest := 0
	minimum := math.Inf(+1)
	for i, p := range ps {
		val := geometry.DistanceSquared(p, q)
		if val < minimum {
			minimum = val
			closest = i
		}
	}
	return closest
}

// clusteringForCentroids takes a set of points and a set of centroids and returns a clustering that
// assigns each point to the closest centroid.
func clusteringForCentroids(points [][2]float64, centroids [][2]float64) []int {
	cl := make([]int, len(points))
	for i, p := range points {
		cl[i] = closest(centroids, p)
	}
	return cl
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
