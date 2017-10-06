package index

import (
	"testing"
)

func TestKDTreePoints(t *testing.T) {
	i := NewKDTree(points)
	testPoints(t, i)
}

func TestKDTreeBoundingBox(t *testing.T) {
	i := NewKDTree(points)
	testBoundingBox(t, i)
}

func TestKDTreeCircle(t *testing.T) {
	i := NewKDTree(points)
	testCircle(t, i)
}
