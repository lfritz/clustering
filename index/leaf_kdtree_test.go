package index

import (
	"testing"
)

func TestLeafKDTreePoints(t *testing.T) {
	i := NewLeafKDTree(points)
	testPoints(t, i)
}

func TestLeafKDTreeBoundingBox(t *testing.T) {
	i := NewLeafKDTree(points)
	testBoundingBox(t, i)
}

func TestLeafKDTreeCircle(t *testing.T) {
	i := NewLeafKDTree(points)
	testCircle(t, i)
}
