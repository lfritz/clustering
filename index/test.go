package index

import (
	"github.com/lfritz/clustering"
	"reflect"
	"sort"
	"testing"
)

var points = []clustering.Point{
	{X: 1, Y: 4},
	{X: 2, Y: 6},
	{X: 3, Y: 3},
	{X: 3, Y: 6},
	{X: 4, Y: 1},
	{X: 4, Y: 6},
	{X: 5, Y: 3},
	{X: 5, Y: 6},
	{X: 8, Y: 6},
}

func testPoints(t *testing.T, i Index) {
	got := i.Points()
	if !reflect.DeepEqual(got, points) {
		t.Errorf("i.Points() returned unexpected result")
	}
}

func testBoundingBox(t *testing.T, i Index) {
	cases := []struct {
		x0, x1, y0, y1 float64
		want           []int
	}{
		{6, 8, 1, 4, []int{}},
		{3, 5, 3, 6, []int{2}},
		{3, 5.1, 3, 6.1, []int{2, 3, 5, 6, 7}},
	}
	for _, c := range cases {
		got := i.BoundingBox(c.x0, c.x1, c.y0, c.y1)
		sort.Ints(got)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("i.BoundingBox(%v, %v, %v, %v) returned %v, want %v",
				c.x0, c.x1, c.y0, c.y1, got, c.want)
		}
	}
}

func testCircle(t *testing.T, i Index) {
	cases := []struct {
		center clustering.Point
		radius float64
		want   []int
	}{
		{clustering.Point{X: 2, Y: 2}, 1, []int{}},
		{clustering.Point{X: 2, Y: 2}, 2, []int{2}},
		{clustering.Point{X: 4, Y: 7}, 3, []int{1, 3, 5, 7}},
	}
	for _, c := range cases {
		got := i.Circle(c.center, c.radius)
		sort.Ints(got)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("i.Circle(%v, %v) returned %v, want %v",
				c.center, c.radius, got, c.want)
		}
	}
}
