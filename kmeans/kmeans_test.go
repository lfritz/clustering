package kmeans

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestRandK(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := 1 + rand.Intn(100)
		k := 1 + rand.Intn(n)

		got := randK(k, n)

		ok := true
		seen := make([]bool, n)
		for _, x := range got {
			if x < 0 || x >= n || seen[x] {
				ok = false
				break
			}
			seen[x] = true
		}
		if !ok {
			t.Errorf("randK(%v, %v) returned %v", k, n, got)
		}
	}
}

var kmeansTestPoints = [][2]float64{
	{1, 8}, {1, 9}, {2, 8},
	{2, 1}, {3, 1}, {3, 2},
	{8, 5}, {8, 7}, {9, 6}, {9, 7},
}

func TestChooseWeightedFor(t *testing.T) {
	weights := []float64{20, 30, 5, 10, 35}
	cases := []struct {
		r    float64
		want int
	}{
		{.51, 2},
		{.30, 1},
		{.90, 4},
		{.55, 3},
		{.00, 0},
	}
	for _, c := range cases {
		got := chooseWeightedFor(weights, c.r)
		if got != c.want {
			t.Errorf("chooseWeightedFor(%v, %v) returned %v, want %v",
				weights, c.r, got, c.want)
		}
	}
}

func TestInitPlusPlus(t *testing.T) {
	k := 3
	got := initPlusPlus(kmeansTestPoints, k)
	ok := true
	if len(got) != k {
		ok = false
	}
	for i, centroid := range got {
		for _, other := range got[:i] {
			if centroid == other {
				ok = false
			}
		}
	}
	if !ok {
		t.Errorf("initPlusPlus(%v, %v) returned %v", kmeansTestPoints, k, got)
	}
}

func TestClusteringForCentroids(t *testing.T) {
	cases := []struct {
		centroids [][2]float64
		want      []int
	}{
		{[][2]float64{{2, 9}, {2, 3}, {8, 6}}, []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2}},
		{[][2]float64{{5, 5}, {5, 6}, {5, 7}}, []int{2, 2, 2, 0, 0, 0, 0, 2, 1, 2}},
		{[][2]float64{{1, 4}, {2, 4}, {3, 4}}, []int{0, 0, 1, 1, 2, 2, 2, 2, 2, 2}},
	}
	for _, c := range cases {
		got, _ := clusteringForCentroids(kmeansTestPoints, c.centroids)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("clusteringForCentroids(%v, %v) returned %v, want %v",
				kmeansTestPoints, c.centroids, got, c.want)
		}
	}
}

func TestCentroidsForClusters(t *testing.T) {
	points := [][2]float64{
		{1, 1}, {2, 8}, {3, 1}, {3, 3}, {3, 6}, {4, 7}, {7, 5}, {8, 4}, {8, 5}, {9, 3},
	}
	k := 3
	cl := []int{2, 0, 2, 1, 0, 0, 1, 1, 1, 1}
	want := [][2]float64{{3, 7}, {7, 4}, {2, 1}}
	got := centroidsForClusters(points, k, cl)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("centroidsForClusters(%v, %v, %v) returned %v, want %v",
			points, k, cl, got, want)
	}
}
