package dbscan

import (
	"github.com/lfritz/clustering/geometry"
	"math"
	"math/rand"
	"reflect"
)

// Kmeans implements the k-mean clustering algorithm
// (https://en.m.wikipedia.org/wiki/K-means_clustering).
func Kmeans(points [][2]float64, k int) []int {
	centroids := initialCentroids(points, k)
	var clustering []int = nil
	for {
		nextClustering := clusteringForCentroids(points, centroids)
		if reflect.DeepEqual(nextClustering, clustering) {
			break
		}
		clustering = nextClustering
		centroids = centroidsForClusters(points, k, clustering)

	}
	return clustering
}

// TODO:
// - try using a struct for intermediate results
// - re-name packages
// - try k-means++
// - try running k-means repeatedly

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

// initialCentroids generates an initial set of centroids for the k-mean algorithm using the Forgy
// method.
func initialCentroids(points [][2]float64, k int) [][2]float64 {
	centroids := make([][2]float64, k)
	for i, x := range randK(k, len(points)) {
		centroids[i] = points[x]
	}
	return centroids
}

func closest(ps [][2]float64, q [2]float64) int {
	closest := 0
	minimumDistance := math.Inf(+1)
	for i, p := range ps {
		distance := geometry.Distance(p, q)
		if distance < minimumDistance {
			minimumDistance = distance
			closest = i
		}
	}
	return closest
}

// clusteringForCentroids takes a set of points and a set of centroids and returns a clustering that
// assigns each point to the closest centroid.
func clusteringForCentroids(points [][2]float64, centroids [][2]float64) []int {
	clustering := make([]int, len(points))
	for i, p := range points {
		clustering[i] = closest(centroids, p)
	}
	return clustering
}

// centroidsForClusters calculates a centroid for each cluster as the mean of its points.
func centroidsForClusters(points [][2]float64, k int, clustering []int) [][2]float64 {
	count := make([]int, k)
	centroids := make([][2]float64, k)
	for i, c := range clustering {
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
