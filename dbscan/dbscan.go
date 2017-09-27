/*
Package dbscan implement the DBSCAN (Density-based spatial clustering of applications with noise)
algorithm.

References:

    https://en.m.wikipedia.org/wiki/DBSCAN
    http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.121.9220

*/
package dbscan

import (
	"github.com/lfritz/clustering"
)

func neighbors(points []clustering.Point, p int, eps float64) []int {
	var result []int
	for i, point := range points {
		if clustering.Distance(points[p], point) < eps {
			result = append(result, i)
		}
	}
	return result
}

// Dbscan applies the DBSCAN algorithm and returns a clustering for points.
func Dbscan(points []clustering.Point, eps float64, minPts int) []int {
	clusterID := Noise + 1
	clustering := make([]int, len(points))
	for p := range points {
		if clustering[p] == Unclassified {
			if expandCluster(points, p, clustering, clusterID, eps, minPts) {
				clusterID++
			}
		}
	}
	return clustering
}

func remove(slice []int, element int) []int {
	i := -1
	for j, x := range slice {
		if x == element {
			i = j
			break
		}
	}
	if i != -1 {
		return append(slice[:i], slice[i+1:]...)
	}
	return slice
}

func expandCluster(points []clustering.Point, p int, clustering []int,
	clusterID int, eps float64, minPts int) bool {
	seeds := neighbors(points, p, eps)
	if len(seeds) < minPts {
		// not a core point
		clustering[p] = Noise
		return false
	}

	for _, q := range seeds {
		clustering[q] = clusterID
	}
	remove(seeds, p)

	for len(seeds) > 0 {
		var q int
		q, seeds = seeds[0], seeds[1:]
		qNeighbors := neighbors(points, q, eps)
		if len(qNeighbors) >= minPts {
			// q is a core point
			for _, r := range qNeighbors {
				if clustering[r] == Unclassified || clustering[r] == Noise {
					if clustering[r] == Unclassified {
						seeds = append(seeds, r)
					}
					clustering[r] = clusterID
				}
			}
		}
	}

	return true
}
