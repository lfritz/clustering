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
	"github.com/lfritz/clustering/index"
)

// Dbscan applies the DBSCAN algorithm and returns a clustering for a set of points.
func Dbscan(i index.Index, eps float64, minPts int) []int {
	clusterID := 0
	points := i.Points()
	cl := make([]int, len(points))
	for i := range cl {
		cl[i] = clustering.Unclassified
	}
	for p := range points {
		if cl[p] == clustering.Unclassified {
			if expandCluster(i, p, cl, clusterID, eps, minPts) {
				clusterID++
			}
		}
	}
	return cl
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

func expandCluster(i index.Index, p int, cl []int, clusterID int, eps float64, minPts int) bool {
	seeds := neighbors(i, p, eps)
	if len(seeds) < minPts {
		// not a core point
		cl[p] = clustering.Noise
		return false
	}

	for _, q := range seeds {
		cl[q] = clusterID
	}
	remove(seeds, p)

	for len(seeds) > 0 {
		var q int
		q, seeds = seeds[0], seeds[1:]
		qNeighbors := neighbors(i, q, eps)
		if len(qNeighbors) >= minPts {
			// q is a core point
			for _, r := range qNeighbors {
				if cl[r] == clustering.Unclassified ||
					cl[r] == clustering.Noise {
					if cl[r] == clustering.Unclassified {
						seeds = append(seeds, r)
					}
					cl[r] = clusterID
				}
			}
		}
	}

	return true
}
