package index

import (
	"reflect"
	"sort"
	"testing"
)

var points = [][2]float64{
	{1, 4},
	{2, 6},
	{3, 3},
	{3, 6},
	{4, 1},
	{4, 6},
	{5, 3},
	{5, 6},
	{8, 6},
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
		center [2]float64
		radius float64
		want   []int
	}{
		{[2]float64{2, 2}, 1, []int{}},
		{[2]float64{2, 2}, 2, []int{2}},
		{[2]float64{4, 7}, 3, []int{1, 3, 5, 7}},
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
