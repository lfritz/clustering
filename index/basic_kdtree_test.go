package index

import (
	"testing"
)

func TestBasicKDTreePoints(t *testing.T) {
	i := NewBasicKDTree(points)
	testPoints(t, i)
}

func TestBasicKDTreeBoundingBox(t *testing.T) {
	i := NewBasicKDTree(points)
	testBoundingBox(t, i)
}

func TestBasicKDTreeCircle(t *testing.T) {
	i := NewBasicKDTree(points)
	testCircle(t, i)
}
