package random

import (
	"testing"
)

func TestChooseWeighted(t *testing.T) {
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
		got := chooseWeighted(weights, c.r)
		if got != c.want {
			t.Errorf("chooseWeighted(%v, %v) returned %v, want %v",
				weights, c.r, got, c.want)
		}
	}
}
