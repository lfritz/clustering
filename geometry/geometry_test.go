package geometry

import (
	"testing"
)

func TestDistanceSquared(t *testing.T) {
	cases := []struct {
		a, b [2]float64
		want float64
	}{
		{[2]float64{1, 2}, [2]float64{1, 2}, 0},
		{[2]float64{1, 2}, [2]float64{3, 2}, 4},
		{[2]float64{1, 2}, [2]float64{1, 4}, 4},
		{[2]float64{1, 2}, [2]float64{4, 6}, 25},
	}
	for _, c := range cases {
		got := DistanceSquared(c.a, c.b)
		if got != c.want {
			t.Errorf("distance(%v, %v) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		a, b [2]float64
		want float64
	}{
		{[2]float64{1, 2}, [2]float64{1, 2}, 0},
		{[2]float64{1, 2}, [2]float64{3, 2}, 2},
		{[2]float64{1, 2}, [2]float64{1, 4}, 2},
		{[2]float64{1, 2}, [2]float64{4, 6}, 5},
	}
	for _, c := range cases {
		got := Distance(c.a, c.b)
		if got != c.want {
			t.Errorf("distance(%v, %v) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
