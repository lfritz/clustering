package dbscan

import (
	"github.com/lfritz/clustering"
	"github.com/lfritz/clustering/index"
	"reflect"
	"testing"
)

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

var testPoints = [][2]float64{
	// cluster a: 3 points
	{1, 8}, {1, 7}, {2, 7},
	// cluster b: 8 points
	{6, 8}, {7, 8},
	{5, 7}, {6, 7}, {7, 7}, {8, 7},
	{6, 6}, {7, 6},
	// cluster c: 8 points
	{2, 3}, {3, 3}, {1, 2}, {2, 2}, {3, 2}, {2, 1}, {3, 1},
	{4, 2}, // border point of both c and d
	// cluster d: 5 points
	{5, 3}, {5, 2}, {6, 2}, {5, 1},
}

var testIndex = index.NewTrivialIndex(testPoints)

func TestDbscan(t *testing.T) {
	expected := []int{
		clustering.Noise, clustering.Noise, clustering.Noise,
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1,
		2, 2, 2, 2, 2,
	}
	cl := Dbscan(testIndex, 1.1, 4)
	if !reflect.DeepEqual(cl, expected) {
		t.Errorf("Dbscan(testIndex, 1.1, 4)\nreturned: %v\nexpected: %v",
			cl, expected)
	}
}

func TestExpandCluster(t *testing.T) {
	// initial clustering: cluster d is marked 2, everything else unclassified
	initialClustering := make([]int, len(testPoints))
	for i := range initialClustering {
		if 19 <= i && i < 23 {
			initialClustering[i] = 0
		} else {
			initialClustering[i] = clustering.Unclassified
		}
	}

	// returns a copy of initialClustering with a range of values changed
	changeClustering := func(start, end, value int) []int {
		cl := make([]int, len(initialClustering))
		copy(cl, initialClustering)
		for i := start; i < end; i++ {
			cl[i] = value
		}
		return cl
	}

	cases := []struct {
		p                  int
		minPts             int
		expectedResult     bool
		expectedClustering []int
	}{
		{1, 3, true, changeClustering(0, 3, 1)},
		{1, 4, false, changeClustering(1, 2, clustering.Noise)},
		{6, 4, true, changeClustering(3, 11, 1)},
		{14, 4, true, changeClustering(11, 19, 1)},
	}

	cl := make([]int, len(initialClustering))
	for _, c := range cases {
		copy(cl, initialClustering)
		result := expandCluster(testIndex, c.p, cl, 1, 1.1, c.minPts)
		if !(result == c.expectedResult &&
			reflect.DeepEqual(cl, c.expectedClustering)) {
			t.Errorf("expandCluster(testIndex, %v, cl, 3, 1.1, %v):\n"+
				"expected: %v, %v\n"+
				"got:      %v, %v",
				c.p, c.minPts,
				c.expectedClustering, c.expectedResult,
				cl, result)
		}
	}
}
