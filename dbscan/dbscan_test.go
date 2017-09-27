package dbscan

import (
	"github.com/lfritz/clustering"
	"reflect"
	"testing"
)

func TestNeighbors(t *testing.T) {
	points := []clustering.Point{
		{X: 1, Y: 16}, // no neighbors
		{X: 1, Y: 4},  // one neighbor
		{X: 1.2, Y: 4.3},
		{X: 12, Y: 16}, // no neighbors (but one point is very close)
		{X: 12.9, Y: 16.5},
		{X: 15, Y: 16}, // five neighbors
		{X: 14.8, Y: 15.9},
		{X: 15.2, Y: 15.7},
		{X: 14.8, Y: 16.5},
		{X: 15.1, Y: 15.8},
		{X: 15.4, Y: 16.2},
	}
	cases := []struct {
		in   int
		want []int
	}{
		{0, []int{0}},
		{1, []int{1, 2}},
		{3, []int{3}},
		{5, []int{5, 6, 7, 8, 9, 10}},
	}
	for _, c := range cases {
		got := neighbors(points, c.in, 1)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("neighbors(points, %v, 1) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		in      []int
		element int
		want    []int
	}{
		{[]int{}, 5, []int{}},
		{[]int{3, 6, 2}, 5, []int{3, 6, 2}},
		{[]int{3, 5, 6, 2}, 5, []int{3, 6, 2}},
		{[]int{3, 5, 6, 2, 5}, 5, []int{3, 6, 2, 5}},
	}
	for _, c := range cases {
		got := remove(c.in, c.element)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("remove(%v, %v) == %v, want %v", c.in, c.element, got, c.want)
		}
	}
}

var points = []clustering.Point{
	// cluster a: 3 points
	{X: 1, Y: 8}, {X: 1, Y: 7}, {X: 2, Y: 7},
	// cluster b: 8 points
	{X: 6, Y: 8}, {X: 7, Y: 8},
	{X: 5, Y: 7}, {X: 6, Y: 7}, {X: 7, Y: 7}, {X: 8, Y: 7},
	{X: 6, Y: 6}, {X: 7, Y: 6},
	// cluster c: 8 points
	{X: 2, Y: 3}, {X: 3, Y: 3}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 3, Y: 2}, {X: 2, Y: 1}, {X: 3, Y: 1},
	{X: 4, Y: 2}, // border point of both c and d
	// cluster d: 5 points
	{X: 5, Y: 3}, {X: 5, Y: 2}, {X: 6, Y: 2}, {X: 5, Y: 1},
}

func TestDbscan(t *testing.T) {
	expectedClustering := []int{
		Noise, Noise, Noise,
		2, 2, 2, 2, 2, 2, 2, 2,
		3, 3, 3, 3, 3, 3, 3,
		4, 4, 4, 4, 4,
	}
	clustering := Dbscan(points, 1.1, 4)
	if !reflect.DeepEqual(clustering, expectedClustering) {
		t.Errorf("Dbscan(points, 1.1, 4)\nreturned: %v\nexpected: %v",
			clustering, expectedClustering)
	}
}

func TestExpandCluster(t *testing.T) {
	// initial clustering: cluster d is marked 2, everything else unclassified
	initialClustering := make([]int, len(points))
	for i := 19; i < 23; i++ {
		initialClustering[i] = 2
	}

	// returns a copy of initialClustering with a range of values changed
	changeClustering := func(start, end, value int) []int {
		clustering := make([]int, len(initialClustering))
		copy(clustering, initialClustering)
		for i := start; i < end; i++ {
			clustering[i] = value
		}
		return clustering
	}

	cases := []struct {
		p                  int
		minPts             int
		expectedResult     bool
		expectedClustering []int
	}{
		{1, 3, true, changeClustering(0, 3, 3)},
		{1, 4, false, changeClustering(1, 2, Noise)},
		{6, 4, true, changeClustering(3, 11, 3)},
		{14, 4, true, changeClustering(11, 19, 3)},
	}

	clustering := make([]int, len(initialClustering))
	for _, c := range cases {
		copy(clustering, initialClustering)
		result := expandCluster(points, c.p, clustering, 3, 1.1, c.minPts)
		if !(result == c.expectedResult &&
			reflect.DeepEqual(clustering, c.expectedClustering)) {
			t.Errorf("expandCluster(points, %v, clustering, 3, 1.1, %v):\n"+
				"expected: %v, %v\n"+
				"got:      %v, %v",
				c.p, c.minPts,
				c.expectedClustering, c.expectedResult,
				clustering, result)
		}
	}
}
