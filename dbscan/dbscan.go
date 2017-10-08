/*
Package dbscan implement the DBSCAN (Density-based spatial clustering of applications with noise)
algorithm.

References:

    https://en.m.wikipedia.org/wiki/DBSCAN
    http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.121.9220

*/
package dbscan

import (
	"github.com/lfritz/clustering/index"
)

// Dbscan applies the DBSCAN algorithm and returns a clustering for a set of points.
func Dbscan(i index.Index, eps float64, minPts int) []int {
	clusterID := Noise + 1
	points := i.Points()
	clustering := make([]int, len(points))
	for p := range points {
		if clustering[p] == Unclassified {
			if expandCluster(i, p, clustering, clusterID, eps, minPts) {
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

func neighbors(i index.Index, p int, eps float64) []int {
	return index.Circle(i, i.Points()[p], eps)
}

func expandCluster(i index.Index, p int, clustering []int,
	clusterID int, eps float64, minPts int) bool {
	seeds := neighbors(i, p, eps)
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
		qNeighbors := neighbors(i, q, eps)
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
