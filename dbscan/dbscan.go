package dbscan

import (
	"math"
	"github.com/lfritz/clustering"
)

const (
	Unclassified = 0
	Noise = 1
)

func distance(a, b clustering.Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func neighbors(points []clustering.Point, p int, eps float64) []int {
	var result []int
	for i, point := range points {
		if distance(points[p], point) < eps {
			result = append(result, i)
		}
	}
	return result
}

func Dbscan(points []clustering.Point, eps float64, minPts int) []int {
	clusterId := Noise + 1
	clustering := make([]int, len(points))
	for p := range points {
		if clustering[p] == Unclassified {
			if expandCluster(points, p, clustering, clusterId, eps, minPts) {
				clusterId++
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
	clusterId int, eps float64, minPts int) bool {
	seeds := neighbors(points, p, eps)
	if len(seeds) < minPts {
		// not a core point
		clustering[p] = Noise
		return false
	}

	for _, q := range seeds {
		clustering[q] = clusterId
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
					clustering[r] = clusterId
				}
			}
		}
	}

	return true
}

