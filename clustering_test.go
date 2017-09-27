package clustering

import (
	"testing"
)

func TestDistance(t *testing.T) {
	cases := []struct {
		a, b Point
		want float64
	}{
		{Point{1, 2}, Point{1, 2}, 0},
		{Point{1, 2}, Point{3, 2}, 2},
		{Point{1, 2}, Point{1, 4}, 2},
		{Point{1, 2}, Point{4, 6}, 5},
	}
	for _, c := range cases {
		got := Distance(c.a, c.b)
		if got != c.want {
			t.Errorf("distance(%v, %v) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
